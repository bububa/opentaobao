package wlbimports

// TransferstoreQueryRequest 结构体
type TransferstoreQueryRequest struct {
	// 运输方式（卡车：TRUCK，快递：EXPRESS）
	TransportType string `json:"transport_type,omitempty" xml:"transport_type,omitempty"`
	// 发件人联系地址ID
	FromId int64 `json:"from_id,omitempty" xml:"from_id,omitempty"`
	// 发件人国家ID
	ToId int64 `json:"to_id,omitempty" xml:"to_id,omitempty"`
}
