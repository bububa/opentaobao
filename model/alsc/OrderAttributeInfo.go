package alsc

// OrderAttributeInfo 
/* model for simplify = false
type OrderAttributeInfo struct {

    // 属性所属组:ORDER 订单类，SUBORDER 子订单类
    
    AttrGroup   string `json:"attr_group,omitempty"`
    

    // 套餐明细
    
    ComboInfoList  struct {
        ComboInfo  []ComboInfo `json:"combo_info,omitempty"`
    } `json:"combo_info_list,omitempty"`
    

    // 单品做法明细
    
    CookingMethodsInfoList  struct {
        AttachInfo  []AttachInfo `json:"attach_info,omitempty"`
    } `json:"cooking_methods_info_list,omitempty"`
    

    // 单品配料明细
    
    IngredientsInfoList  struct {
        AttachInfo  []AttachInfo `json:"attach_info,omitempty"`
    } `json:"ingredients_info_list,omitempty"`
    

    // 业务ID,取值与属性所属组对应
    
    OutBizId   string `json:"out_biz_id,omitempty"`
    

    // 业务方订单号
    
    OutOrderNo   string `json:"out_order_no,omitempty"`
    

    // 服务费明细
    
    ServiceFeeInfoList  struct {
        ServiceFeeInfo  []ServiceFeeInfo `json:"service_fee_info,omitempty"`
    } `json:"service_fee_info_list,omitempty"`
    

    // 其他附加费，如餐盒等等
    
    OtherInfoList  struct {
        AttachInfo  []AttachInfo `json:"attach_info,omitempty"`
    } `json:"other_info_list,omitempty"`
    

}
*/

// OrderAttributeInfo 
type OrderAttributeInfo struct {

    // 属性所属组:ORDER 订单类，SUBORDER 子订单类
    AttrGroup   string `json:"attr_group,omitempty"`

    // 套餐明细
    ComboInfoList   []ComboInfo `json:"combo_info_list,omitempty"`

    // 单品做法明细
    CookingMethodsInfoList   []AttachInfo `json:"cooking_methods_info_list,omitempty"`

    // 单品配料明细
    IngredientsInfoList   []AttachInfo `json:"ingredients_info_list,omitempty"`

    // 业务ID,取值与属性所属组对应
    OutBizId   string `json:"out_biz_id,omitempty"`

    // 业务方订单号
    OutOrderNo   string `json:"out_order_no,omitempty"`

    // 服务费明细
    ServiceFeeInfoList   []ServiceFeeInfo `json:"service_fee_info_list,omitempty"`

    // 其他附加费，如餐盒等等
    OtherInfoList   []AttachInfo `json:"other_info_list,omitempty"`

}
