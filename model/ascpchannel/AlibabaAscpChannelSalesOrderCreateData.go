package ascpchannel

// AlibabaAscpChannelSalesOrderCreateData 
type AlibabaAscpChannelSalesOrderCreateData struct {
    // 子单列表
    SubOrderList   []Suborders `json:"sub_order_list,omitempty" xml:"sub_order_list>suborders,omitempty"`
    // 渠道订单号
    SaleOrderNo   string `json:"sale_order_no,omitempty" xml:"sale_order_no,omitempty"`
}
