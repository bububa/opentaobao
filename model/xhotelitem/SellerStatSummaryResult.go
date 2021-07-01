package xhotelitem

// SellerStatSummaryResult 结构体
type SellerStatSummaryResult struct {
	// rate平均分
	AvgRateScore string `json:"avg_rate_score,omitempty" xml:"avg_rate_score,omitempty"`
	// 曝光率
	ExposedPercent string `json:"exposed_percent,omitempty" xml:"exposed_percent,omitempty"`
	// supplier参数
	SupplierParam string `json:"supplier_param,omitempty" xml:"supplier_param,omitempty"`
	// 标准酒店维度曝光总数
	ShidTotalAmount string `json:"shid_total_amount,omitempty" xml:"shid_total_amount,omitempty"`
	// hid参数
	HidParam string `json:"hid_param,omitempty" xml:"hid_param,omitempty"`
	// rate最低分
	MinRateScore string `json:"min_rate_score,omitempty" xml:"min_rate_score,omitempty"`
	// 不可售情况
	UnsaleReseasonInfo string `json:"unsale_reseason_info,omitempty" xml:"unsale_reseason_info,omitempty"`
	// rate最高分
	MaxRateScore string `json:"max_rate_score,omitempty" xml:"max_rate_score,omitempty"`
	// 选品情况
	SelectionMessageInfo string `json:"selection_message_info,omitempty" xml:"selection_message_info,omitempty"`
	// 日期
	DateParam string `json:"date_param,omitempty" xml:"date_param,omitempty"`
	// 商品总数
	TotalAmount string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// vendor参数
	VendorParam string `json:"vendor_param,omitempty" xml:"vendor_param,omitempty"`
	// 曝光总数
	ExposedAmount string `json:"exposed_amount,omitempty" xml:"exposed_amount,omitempty"`
	// 选品情况
	SelectionMessageInfoJson string `json:"selection_message_info_json,omitempty" xml:"selection_message_info_json,omitempty"`
	// 不可售情况
	UnsaleReasonInfoJson string `json:"unsale_reason_info_json,omitempty" xml:"unsale_reason_info_json,omitempty"`
	// sellerId参数
	SellerIdParam string `json:"seller_id_param,omitempty" xml:"seller_id_param,omitempty"`
	// 可售商品数量
	CanSaleAmount string `json:"can_sale_amount,omitempty" xml:"can_sale_amount,omitempty"`
	// 选品保留商品数量
	SelectedAmount string `json:"selected_amount,omitempty" xml:"selected_amount,omitempty"`
}
