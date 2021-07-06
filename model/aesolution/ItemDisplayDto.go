package aesolution

// ItemDisplayDto 结构体
type ItemDisplayDto struct {
	// product offline time
	WsOfflineDate string `json:"ws_offline_date,omitempty" xml:"ws_offline_date,omitempty"`
	// product offline reason
	WsDisplay string `json:"ws_display,omitempty" xml:"ws_display,omitempty"`
	// product tite
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// product src
	Src string `json:"src,omitempty" xml:"src,omitempty"`
	// min price among all skus of the product
	ProductMinPrice string `json:"product_min_price,omitempty" xml:"product_min_price,omitempty"`
	// max price among all skus of the product
	ProductMaxPrice string `json:"product_max_price,omitempty" xml:"product_max_price,omitempty"`
	// seller login id
	OwnerMemberId string `json:"owner_member_id,omitempty" xml:"owner_member_id,omitempty"`
	// product image urls
	ImageURLs string `json:"image_u_r_ls,omitempty" xml:"image_u_r_ls,omitempty"`
	// time that product was modifed
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// time that product was created
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// currency code
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// Coupon start date, GMT+8
	CouponStartDate string `json:"coupon_start_date,omitempty" xml:"coupon_start_date,omitempty"`
	// Coupon end date, GMT+8
	CouponEndDate string `json:"coupon_end_date,omitempty" xml:"coupon_end_date,omitempty"`
	// product id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// seller member seq
	OwnerMemberSeq int64 `json:"owner_member_seq,omitempty" xml:"owner_member_seq,omitempty"`
	// group id
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// freight template id
	FreightTemplateId int64 `json:"freight_template_id,omitempty" xml:"freight_template_id,omitempty"`
}
