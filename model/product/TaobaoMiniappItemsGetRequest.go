package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取商品信息 API请求
taobao.miniapp.items.get

获取商品公开属性，只允许在商家应用环境中使用
*/
type TaobaoMiniappItemsGetRequest struct {
    model.Params
    // 商品数字id列表，多个num_iid用逗号隔开，一次不超过50个。
    numIids   []string
    // 需要返回的商品对象字段。可选值：Item商品结构体中字段均可返回(其中item_weight,item_size,sold_quantity暂未返回)；多个字段用“,”分隔。如果想返回整个子对象，fields设置相应字段，如itemimg；如果想返回子对象里面的某个字段，那字段设为某个值，如itemimg.url。
    fields   []string
}

// 初始化TaobaoMiniappItemsGetRequest对象
func NewTaobaoMiniappItemsGetRequest() *TaobaoMiniappItemsGetRequest{
    return &TaobaoMiniappItemsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappItemsGetRequest) GetApiMethodName() string {
    return "taobao.miniapp.items.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappItemsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NumIids Setter
// 商品数字id列表，多个num_iid用逗号隔开，一次不超过50个。
func (r *TaobaoMiniappItemsGetRequest) SetNumIids(numIids []string) error {
    r.numIids = numIids
    r.Set("num_iids", numIids)
    return nil
}

// NumIids Getter
func (r TaobaoMiniappItemsGetRequest) GetNumIids() []string {
    return r.numIids
}
// Fields Setter
// 需要返回的商品对象字段。可选值：Item商品结构体中字段均可返回(其中item_weight,item_size,sold_quantity暂未返回)；多个字段用“,”分隔。如果想返回整个子对象，fields设置相应字段，如itemimg；如果想返回子对象里面的某个字段，那字段设为某个值，如itemimg.url。
func (r *TaobaoMiniappItemsGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoMiniappItemsGetRequest) GetFields() []string {
    return r.fields
}
