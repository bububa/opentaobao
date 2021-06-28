package promotion

// CouponTemplateOperateResponse 
/* model for simplify = false
type CouponTemplateOperateResponse struct {

    // 模板表主键
    
    Id   int64 `json:"id,omitempty"`
    

    // ump模板ID
    
    SourceId   int64 `json:"source_id,omitempty"`
    

    // 创建结果
    
    FailElements  struct {
        FailElement  []FailElement `json:"fail_element,omitempty"`
    } `json:"fail_elements,omitempty"`
    

    // 分页信息
    
    PageInfo  *struct {
        PageInfo  *PageInfo `json:"page_info,omitempty"`
    } `json:"page_info,omitempty"`
    

    // 券圈品设置
    
    PromActSkuList  struct {
        PromActSku  []PromActSku `json:"prom_act_sku,omitempty"`
    } `json:"prom_act_sku_list,omitempty"`
    

    // 匿名码code
    
    MaCode   string `json:"ma_code,omitempty"`
    

    // 券实例id
    
    VoucherId   string `json:"voucher_id,omitempty"`
    

    // 券模版
    
    CouponTemplate  *struct {
        CouponTemplate  *CouponTemplate `json:"coupon_template,omitempty"`
    } `json:"coupon_template,omitempty"`
    

}
*/

// CouponTemplateOperateResponse 
type CouponTemplateOperateResponse struct {

    // 模板表主键
    Id   int64 `json:"id,omitempty"`

    // ump模板ID
    SourceId   int64 `json:"source_id,omitempty"`

    // 创建结果
    FailElements   []FailElement `json:"fail_elements,omitempty"`

    // 分页信息
    PageInfo   *PageInfo `json:"page_info,omitempty"`

    // 券圈品设置
    PromActSkuList   []PromActSku `json:"prom_act_sku_list,omitempty"`

    // 匿名码code
    MaCode   string `json:"ma_code,omitempty"`

    // 券实例id
    VoucherId   string `json:"voucher_id,omitempty"`

    // 券模版
    CouponTemplate   *CouponTemplate `json:"coupon_template,omitempty"`

}
