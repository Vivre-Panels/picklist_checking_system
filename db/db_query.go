package database

import (
	"log"

	"picklist_checking_system/models"
)

func SaveWebhookLog(logEntry models.SalesOrder) error {
	query := `
	IF NOT EXISTS (SELECT 1 FROM Sales_Order WHERE salesorder_id = @p1 and item_id = @p2)
		INSERT INTO Sales_Order (
			salesorder_id,
			item_id,
			name,
			quantity,
			rate,
			item_sub_total_formatted,
			unit,
			received_at,
		) VALUES (
			@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8
		);
	ELSE
		UPDATE job_description
		SET formatted_jd = @p2,
			questions = @p3,
			job_title = @p4,
			job_status = @p5,
			work_experience = @p6,
			updated_at = @p7
		WHERE record_id = @p1;
	`
	


	_, err := db.Exec(query,
		logEntry.SalesorderID,
		logEntry.ItemID,
		logEntry.ItemName,
		logEntry.Quantity,
		logEntry.Rate,
		logEntry.Unit,
		logEntry.ItemSubTotalFormatted,
		logEntry.ReceivedAT,
		logEntry.UpdateAT,
	)
	if err != nil {
		log.Println("Webhook log insert error:", err)
	}
	return err
}
