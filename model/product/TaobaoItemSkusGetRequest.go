package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据商品ID列表获取SKU信息 API请求
taobao.item.skus.get

* 获取多个商品下的所有sku
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoItemSkusGetAPIRequest struct {
    model.Params
    // 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。
    _fields   []string
    // sku所属商品数字id，必选。num_iid个数不能超过40个
    _numIids   string
}

// 初始化TaobaoItemSkusGetAPIRequest对象
func NewTaobaoItemSkusGetRequest() *TaobaoItemSkusGetAPIRequest{
    return &TaobaoItemSkusGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemSkusGetAPIRequest) GetApiMethodName() string {
    return "taobao.item.skus.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemSkusGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。
func (r *TaobaoItemSkusGetAPIRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoItemSkusGetAPIRequest) GetFields() []string {
    return r._fields
}
// NumIids Setter
// sku所属商品数字id，必选。num_iid个数不能超过40个
func (r *TaobaoItemSkusGetAPIRequest) SetNumIids(_numIids string) error {
    r._numIids = _numIids
    r.Set("num_iids", _numIids)
    return nil
}

// NumIids Getter
func (r TaobaoItemSkusGetAPIRequest) GetNumIids() string {
    return r._numIids
}
