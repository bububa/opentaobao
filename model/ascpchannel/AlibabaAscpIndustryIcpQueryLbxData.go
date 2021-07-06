package ascpchannel

// AlibabaAscpIndustryIcpQueryLbxData 结构体
type AlibabaAscpIndustryIcpQueryLbxData struct {
	// 单信息
	TransferDetailList []Transferdetaildtolist `json:"transfer_detail_list,omitempty" xml:"transfer_detail_list>transferdetaildtolist,omitempty"`
	// 外部icps单号
	OutBizCode string `json:"out_biz_code,omitempty" xml:"out_biz_code,omitempty"`
	// 调拨状态
	TransferOrderStatus string `json:"transfer_order_status,omitempty" xml:"transfer_order_status,omitempty"`
}
