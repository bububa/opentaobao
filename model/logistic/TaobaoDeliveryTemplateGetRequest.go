package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户指定运费模板信息 API请求
taobao.delivery.template.get

获取用户指定运费模板信息
*/
type TaobaoDeliveryTemplateGetRequest struct {
    model.Params
    // 需要查询的运费模板ID列表
    templateIds   []string
    // 需返回的字段列表。 <br/> <br/><B><br/>可选值:示例中的值;各字段之间用","隔开<br/></B><br/><br/><br/> <br/><font color=blue><br/>template_id：查询模板ID <br/> <br/>template_name:查询模板名称 <br/> <br/>assumer：查询服务承担放 <br/> <br/>valuation:查询计价规则 <br/> <br/>supports:查询增值服务列表 <br/> <br/>created:查询模板创建时间 <br/> <br/>modified:查询修改时间<br/> <br/>consign_area_id:运费模板上设置的卖家发货地址最小级别ID<br/> <br/>address:卖家地址信息<br/></font><br/><br/><br/> <br/><font color=bule><br/>query_cod:查询货到付款运费信息<br/> <br/>query_post:查询平邮运费信息 <br/> <br/>query_express:查询快递公司运费信息 <br/> <br/>query_ems:查询EMS运费信息<br/> <br/>query_bzsd:查询普通保障速递运费信息<br/> <br/>query_wlb:查询物流宝保障速递运费信息<br/> <br/>query_furniture:查询家装物流运费信息<br/><font color=blue><br/><br/><br/><br/><font color=red>不能有空格</font>
    fields   []string
    // 在没有登录的情况下根据淘宝用户昵称查询指定的模板
    userNick   string
}

// 初始化TaobaoDeliveryTemplateGetRequest对象
func NewTaobaoDeliveryTemplateGetRequest() *TaobaoDeliveryTemplateGetRequest{
    return &TaobaoDeliveryTemplateGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDeliveryTemplateGetRequest) GetApiMethodName() string {
    return "taobao.delivery.template.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDeliveryTemplateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TemplateIds Setter
// 需要查询的运费模板ID列表
func (r *TaobaoDeliveryTemplateGetRequest) SetTemplateIds(templateIds []string) error {
    r.templateIds = templateIds
    r.Set("template_ids", templateIds)
    return nil
}

// TemplateIds Getter
func (r TaobaoDeliveryTemplateGetRequest) GetTemplateIds() []string {
    return r.templateIds
}
// Fields Setter
// 需返回的字段列表。 <br/> <br/><B><br/>可选值:示例中的值;各字段之间用","隔开<br/></B><br/><br/><br/> <br/><font color=blue><br/>template_id：查询模板ID <br/> <br/>template_name:查询模板名称 <br/> <br/>assumer：查询服务承担放 <br/> <br/>valuation:查询计价规则 <br/> <br/>supports:查询增值服务列表 <br/> <br/>created:查询模板创建时间 <br/> <br/>modified:查询修改时间<br/> <br/>consign_area_id:运费模板上设置的卖家发货地址最小级别ID<br/> <br/>address:卖家地址信息<br/></font><br/><br/><br/> <br/><font color=bule><br/>query_cod:查询货到付款运费信息<br/> <br/>query_post:查询平邮运费信息 <br/> <br/>query_express:查询快递公司运费信息 <br/> <br/>query_ems:查询EMS运费信息<br/> <br/>query_bzsd:查询普通保障速递运费信息<br/> <br/>query_wlb:查询物流宝保障速递运费信息<br/> <br/>query_furniture:查询家装物流运费信息<br/><font color=blue><br/><br/><br/><br/><font color=red>不能有空格</font>
func (r *TaobaoDeliveryTemplateGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoDeliveryTemplateGetRequest) GetFields() []string {
    return r.fields
}
// UserNick Setter
// 在没有登录的情况下根据淘宝用户昵称查询指定的模板
func (r *TaobaoDeliveryTemplateGetRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

// UserNick Getter
func (r TaobaoDeliveryTemplateGetRequest) GetUserNick() string {
    return r.userNick
}
