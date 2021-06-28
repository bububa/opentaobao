package wdk

// FinanceOrderDetailResponse 
/* model for simplify = false
type FinanceOrderDetailResponse struct {

    // 分页信息
    
    Pagination  *struct {
        Pagination  *Pagination `json:"pagination,omitempty"`
    } `json:"pagination,omitempty"`
    

    // 财务订单信息
    
    FinanceOrderDetails  struct {
        FinanceOrderDetail  []FinanceOrderDetail `json:"finance_order_detail,omitempty"`
    } `json:"finance_order_details,omitempty"`
    

}
*/

// FinanceOrderDetailResponse 
type FinanceOrderDetailResponse struct {

    // 分页信息
    Pagination   *Pagination `json:"pagination,omitempty"`

    // 财务订单信息
    FinanceOrderDetails   []FinanceOrderDetail `json:"finance_order_details,omitempty"`

}
