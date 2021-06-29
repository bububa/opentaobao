package retail

// CreateOrderRequest 
type CreateOrderRequest struct {
    // 门店编码
    StoreCode   string `json:"store_code,omitempty" xml:"store_code,omitempty"`
    // 订单列表
    Orders   []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
    // 用户手机号码
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
}
