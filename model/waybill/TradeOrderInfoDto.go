package waybill

// TradeOrderInfoDTO 
type TradeOrderInfoDTO struct {
    // 物流服务值（详见https://support-cnkuaidi.taobao.com/doc.htm#?docId=106156&docType=1，如无特殊服务请置空）
    LogisticsServices   string `json:"logistics_services,omitempty" xml:"logistics_services,omitempty"`
    // <a href="http://open.taobao.com/docs/doc.htm?docType=1&articleId=105086&treeId=17&platformId=17#6">请求ID</a>
    ObjectId   string `json:"object_id,omitempty" xml:"object_id,omitempty"`
    // 订单信息
    OrderInfo   *OrderInfoDTO `json:"order_info,omitempty" xml:"order_info,omitempty"`
    // 包裹信息
    PackageInfo   *PackageInfoDTO `json:"package_info,omitempty" xml:"package_info,omitempty"`
    // 收件人信息
    Recipient   *RecipientInfoDTO `json:"recipient,omitempty" xml:"recipient,omitempty"`
    // 云打印标准模板URL（组装云打印结果使用，值格式http://cloudprint.cainiao.com/template/standard/${模板ID}）
    TemplateUrl   string `json:"template_url,omitempty" xml:"template_url,omitempty"`
    // 使用者ID（使用电子面单账号的实际商家ID，如存在一个电子面单账号多个店铺使用时，请传入店铺的商家ID）
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
