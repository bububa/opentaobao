package aesolution

// ItemListQuery 结构体
type ItemListQuery struct {
	// Current page of products to be needed. The default page is page 1.
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// Product Ids which needs to be excluded
	ExceptedProductIds []int64 `json:"excepted_product_ids,omitempty" xml:"excepted_product_ids>int64,omitempty"`
	// Search field by expiration date. For example, if the value for expiration date is 3, it means to query products to be offline within 3 days.
	OffLineTime int64 `json:"off_line_time,omitempty" xml:"off_line_time,omitempty"`
	// Login ID of product owner
	OwnerMemberId string `json:"owner_member_id,omitempty" xml:"owner_member_id,omitempty"`
	// Number of products to be queried at each page. The input value must be less than 100, the default value of which is 20.
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// product id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// onSelling	Product operation status. Currently, it is divided into 4 types with the following input parameters respectively: onSelling; offline; auditing; and editingRequired.
	ProductStatusType string `json:"product_status_type,omitempty" xml:"product_status_type,omitempty"`
	// Fuzzy search field by product subject. It only supports half-width numbers in English with a length not more than 128.
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// Reasons for product offline: expire_offline; user_offline; violate_offline; punish_offline; and degrade_offline.
	WsDisplay string `json:"ws_display,omitempty" xml:"ws_display,omitempty"`
	// Whether having national quotation. "y" for yes, "n" for no.
	HaveNationalQuote string `json:"have_national_quote,omitempty" xml:"have_national_quote,omitempty"`
	// Search field by product groups. Enter product group id (groupId).
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// Search for products created after a specific time, format: yyyy-MM-dd hh:mm:ss
	GmtCreateStart string `json:"gmt_create_start,omitempty" xml:"gmt_create_start,omitempty"`
	// Search for products created before a specific time，yyyy-MM-dd hh:mm:ss
	GmtCreateEnd string `json:"gmt_create_end,omitempty" xml:"gmt_create_end,omitempty"`
	// Search for product modified after a specific time，yyyy-MM-dd hh:mm:ss
	GmtModifiedStart string `json:"gmt_modified_start,omitempty" xml:"gmt_modified_start,omitempty"`
	// Search for products modified before a specific time，yyyy-MM-dd hh:mm:ss
	GmtModifiedEnd string `json:"gmt_modified_end,omitempty" xml:"gmt_modified_end,omitempty"`
	// merchant sku code
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
}
