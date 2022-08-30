package waybill

// WaybillBranchAccount 结构体
type WaybillBranchAccount struct {
	// 当前网点下的发货地址
	ShippAddressCols []AddressDto `json:"shipp_address_cols,omitempty" xml:"shipp_address_cols>address_dto,omitempty"`
	// 可用的服务信息列表
	ServiceInfoCols []ServiceInfoDto `json:"service_info_cols,omitempty" xml:"service_info_cols>service_info_dto,omitempty"`
	// 月结卡号列表
	CustomerCodeList []string `json:"customer_code_list,omitempty" xml:"customer_code_list>string,omitempty"`
	// 网点Code
	BranchCode string `json:"branch_code,omitempty" xml:"branch_code,omitempty"`
	// 网点名称
	BranchName string `json:"branch_name,omitempty" xml:"branch_name,omitempty"`
	// 号段信息
	SegmentCode string `json:"segment_code,omitempty" xml:"segment_code,omitempty"`
	// 品牌code
	BrandCode string `json:"brand_code,omitempty" xml:"brand_code,omitempty"`
	// 月结卡号map，key为shipp_address_cols.waybill_address_id,value为月结卡号。jsonString
	CustomerCodeMap string `json:"customer_code_map,omitempty" xml:"customer_code_map,omitempty"`
	// 已用面单数量
	AllocatedQuantity int64 `json:"allocated_quantity,omitempty" xml:"allocated_quantity,omitempty"`
	// 网点状态
	BranchStatus int64 `json:"branch_status,omitempty" xml:"branch_status,omitempty"`
	// 取消的面单总数
	CancelQuantity int64 `json:"cancel_quantity,omitempty" xml:"cancel_quantity,omitempty"`
	// 已经打印的面单总数
	PrintQuantity int64 `json:"print_quantity,omitempty" xml:"print_quantity,omitempty"`
	// 电子面单余额数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 商家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}
