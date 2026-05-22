package service

import (
	"log"
	"picklist_checking_system/models"
)

func HandleWebhook(payload map[string]interface{}) {
	// 1. Extract the top-level salesorder object
	salesorder, ok := payload["salesorder"].(map[string]interface{})
	if !ok {
		log.Println("Error: top-level 'salesorder' object not found")
		return
	}

	// 2. Extract salesorder_id from the nested object
	salesOrderID, ok := salesorder["salesorder_id"].(string)
	if !ok {
		log.Println("Warning: salesorder_id not found or not a string")
	}
	log.Printf("Extracted Sales Order ID: %s\n", salesOrderID)

	custom_field_hash := salesorder["custom_field_hash"]
	log.Printf("Extracted custom_field_hash: %s\n", custom_field_hash)

	// 3. Extract cf_creator_ops_id from custom_field_hash
	var creatorOpsID string
	if customFields, ok := salesorder["custom_field_hash"].(map[string]interface{}); ok {
		creatorOpsID, _ = customFields["cf_creator_ops_id"].(string)
	} else {
		log.Println("Warning: custom_field_hash not found")
	}
	log.Printf("Extracted Creator Ops ID: %s\n", creatorOpsID)

	// 3. Extract the line_items slice from the nested object
	lineItemsRaw, ok := salesorder["line_items"].([]interface{})
	if !ok {
		log.Println("Error: line_items not found or not an array")
		return
	}

	// 4. Loop through each item and extract fields
	for i, itemRaw := range lineItemsRaw {
		item, ok := itemRaw.(map[string]interface{})
		if !ok {
			log.Printf("Skipping item at index %d: invalid format\n", i)
			continue
		}

		itemID, _ := item["item_id"].(string)
		name, _ := item["name"].(string)
		quantity, _ := item["quantity"].(float64)
		rate, _ := item["rate"].(float64)
		subTotalFormatted, _ := item["item_sub_total_formatted"].(string)
		unit, _ := item["unit"].(string)

		log.Printf("Item %d -> SalesOrderID: %s, ID: %s, Name: %s, Qty: %.0f, Rate: %.2f, Total: %s, Unit: %s\n",
			i+1, salesOrderID, itemID, name, quantity, rate, subTotalFormatted, unit)
	}
}

func BuildCreatorPayload(
	salesorder map[string]interface{},
	existingMappings map[string]string,
) models.CreatorPayload {

	var payload models.CreatorPayload

	payload.Result.Message = true

	lineItemsRaw := salesorder["line_items"].([]interface{})

	for _, itemRaw := range lineItemsRaw {

		item := itemRaw.(map[string]interface{})

		itemID, _ := item["item_id"].(string)
		name, _ := item["name"].(string)

		qty, _ := item["quantity"].(float64)
		rate, _ := item["rate"].(float64)
		unit, _ := item["unit"].(string)

		amount := qty * rate

		row := models.CreatorSubformRow{
			ProductUniqueCode:   name,
			UsageUnit:           unit,
			Rate:                rate,
			Amount:              amount,
			Qty:                 qty,
			TransferredQuantity: qty,
		}

		// IMPORTANT
		// if mapping exists -> UPDATE existing row
		if creatorSubformID, exists := existingMappings[itemID]; exists {
			row.ID = creatorSubformID
		}

		payload.Data.SubForm = append(payload.Data.SubForm, row)
	}

	return payload
}
