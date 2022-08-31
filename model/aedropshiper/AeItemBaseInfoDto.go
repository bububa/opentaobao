package aedropshiper

// AeItemBaseInfoDto 结构体
type AeItemBaseInfoDto struct {
	// The title of the product
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// The currency unit of the commodity. U.S. Dollar: USD, Ruble: RUB
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// Product status
	ProductStatusType string `json:"product_status_type,omitempty" xml:"product_status_type,omitempty"`
	// Reasons for removal of goods
	WsDisplay string `json:"ws_display,omitempty" xml:"ws_display,omitempty"`
	// The date the product was removed from the shelf
	WsOfflineDate string `json:"ws_offline_date,omitempty" xml:"ws_offline_date,omitempty"`
	// Commodity creation time
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// Change the time
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// Evaluation number
	EvaluationCount string `json:"evaluation_count,omitempty" xml:"evaluation_count,omitempty"`
	// Average rating stars, 1-5 stars
	AvgEvaluationRating string `json:"avg_evaluation_rating,omitempty" xml:"avg_evaluation_rating,omitempty"`
	// Commodity detailed description
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	// Mobile detailed description
	MobileDetail string `json:"mobile_detail,omitempty" xml:"mobile_detail,omitempty"`
	// Item ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// ID of the category of the product
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// Seller&#39;s master account ID
	OwnerMemberSeqLong int64 `json:"owner_member_seq_long,omitempty" xml:"owner_member_seq_long,omitempty"`
}
