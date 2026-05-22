package database

import "log"

func enableAutoCreateTables() {
	createSalesorderSubformMapping := `IF OBJECT_ID(N'dbo.sales_order_subform_mapping', N'U') IS NULL
BEGIN
    CREATE TABLE sales_order_subform_mapping (
        id BIGINT AUTO_INCREMENT PRIMARY KEY,

        sales_order_id VARCHAR(255),
        zoho_item_id VARCHAR(255),

        creator_parent_record_id VARCHAR(255),
        creator_subform_row_id VARCHAR(255),

        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
END`

	execSQL(createSalesorderSubformMapping)

}

func execSQL(query string) {
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Error creating table: ", err)
	}
}
