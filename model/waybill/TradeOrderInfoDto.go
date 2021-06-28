package waybill

// TradeOrderInfoDto 
/* model for simplify = false
type TradeOrderInfoDto struct {

    // 物流服务值（详见https://support-cnkuaidi.taobao.com/doc.htm#?docId=106156&docType=1，如无特殊服务请置空）
    
    LogisticsServices   string `json:"logistics_services,omitempty"`
    

    // <a href="http://open.taobao.com/docs/doc.htm?docType=1&articleId=105086&treeId=17&platformId=17#6">请求ID</a>
    
    ObjectId   string `json:"object_id,omitempty"`
    

    // 订单信息
    
    OrderInfo  *struct {
        OrderInfoDto  *OrderInfoDto `json:"order_info_dto,omitempty"`
    } `json:"order_info,omitempty"`
    

    // 包裹信息
    
    PackageInfo  *struct {
        PackageInfoDto  *PackageInfoDto `json:"package_info_dto,omitempty"`
    } `json:"package_info,omitempty"`
    

    // 收件人信息
    
    Recipient  *struct {
        RecipientInfoDto  *RecipientInfoDto `json:"recipient_info_dto,omitempty"`
    } `json:"recipient,omitempty"`
    

    // 云打印标准模板URL（组装云打印结果使用，值格式http://cloudprint.cainiao.com/template/standard/${模板ID}）
    
    TemplateUrl   string `json:"template_url,omitempty"`
    

    // 使用者ID（使用电子面单账号的实际商家ID，如存在一个电子面单账号多个店铺使用时，请传入店铺的商家ID）
    
    UserId   int64 `json:"user_id,omitempty"`
    

}
*/

// TradeOrderInfoDto 
type TradeOrderInfoDto struct {

    // 物流服务值（详见https://support-cnkuaidi.taobao.com/doc.htm#?docId=106156&docType=1，如无特殊服务请置空）
    LogisticsServices   string `json:"logistics_services,omitempty"`

    // <a href="http://open.taobao.com/docs/doc.htm?docType=1&articleId=105086&treeId=17&platformId=17#6">请求ID</a>
    ObjectId   string `json:"object_id,omitempty"`

    // 订单信息
    OrderInfo   *OrderInfoDto `json:"order_info,omitempty"`

    // 包裹信息
    PackageInfo   *PackageInfoDto `json:"package_info,omitempty"`

    // 收件人信息
    Recipient   *RecipientInfoDto `json:"recipient,omitempty"`

    // 云打印标准模板URL（组装云打印结果使用，值格式http://cloudprint.cainiao.com/template/standard/${模板ID}）
    TemplateUrl   string `json:"template_url,omitempty"`

    // 使用者ID（使用电子面单账号的实际商家ID，如存在一个电子面单账号多个店铺使用时，请传入店铺的商家ID）
    UserId   int64 `json:"user_id,omitempty"`

}
