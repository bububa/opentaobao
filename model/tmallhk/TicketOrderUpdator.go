package tmallhk

// TicketOrderUpdator 
/* model for simplify = false
type TicketOrderUpdator struct {

    // 买手姓名
    
    AgentName   string `json:"agent_name,omitempty"`
    

    // 买手护照过期时间
    
    AgentPassportExpDate   string `json:"agent_passport_exp_date,omitempty"`
    

    // 买手护照截图
    
    AgentPassportsArrList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"agent_passports_arr_list,omitempty"`
    

    // 买手付款时间
    
    AgentPayTime   string `json:"agent_pay_time,omitempty"`
    

    // 子订单号
    
    BizOrderId   int64 `json:"biz_order_id,omitempty"`
    

    // 品牌名
    
    BrandName   string `json:"brand_name,omitempty"`
    

    // 是否锁扣，1是0否
    
    Locker   int64 `json:"locker,omitempty"`
    

    // 锁扣照片
    
    LockerPicturesArrList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"locker_pictures_arr_list,omitempty"`
    

    // 银行付款记录
    
    PaymentRecordsArrList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"payment_records_arr_list,omitempty"`
    

    // 购买地
    
    PurchasedPlace   string `json:"purchased_place,omitempty"`
    

    // 购买地照片
    
    PurchasedPlacePicturesArrList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"purchased_place_pictures_arr_list,omitempty"`
    

    // 小票截图
    
    TicketsArrList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"tickets_arr_list,omitempty"`
    

}
*/

// TicketOrderUpdator 
type TicketOrderUpdator struct {

    // 买手姓名
    AgentName   string `json:"agent_name,omitempty"`

    // 买手护照过期时间
    AgentPassportExpDate   string `json:"agent_passport_exp_date,omitempty"`

    // 买手护照截图
    AgentPassportsArrList   []string `json:"agent_passports_arr_list,omitempty"`

    // 买手付款时间
    AgentPayTime   string `json:"agent_pay_time,omitempty"`

    // 子订单号
    BizOrderId   int64 `json:"biz_order_id,omitempty"`

    // 品牌名
    BrandName   string `json:"brand_name,omitempty"`

    // 是否锁扣，1是0否
    Locker   int64 `json:"locker,omitempty"`

    // 锁扣照片
    LockerPicturesArrList   []string `json:"locker_pictures_arr_list,omitempty"`

    // 银行付款记录
    PaymentRecordsArrList   []string `json:"payment_records_arr_list,omitempty"`

    // 购买地
    PurchasedPlace   string `json:"purchased_place,omitempty"`

    // 购买地照片
    PurchasedPlacePicturesArrList   []string `json:"purchased_place_pictures_arr_list,omitempty"`

    // 小票截图
    TicketsArrList   []string `json:"tickets_arr_list,omitempty"`

}
