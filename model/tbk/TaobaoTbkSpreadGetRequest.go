package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-长链转短链 API请求
taobao.tbk.spread.get

输入一个原始的链接，转换得到指定的传播方式，如二维码，淘口令，短连接；
现阶段只支持短连接。
*/
type TaobaoTbkSpreadGetRequest struct {
    model.Params
    // 请求列表，内部包含多个url
    requests   []TbkSpreadRequest
}

// 初始化TaobaoTbkSpreadGetRequest对象
func NewTaobaoTbkSpreadGetRequest() *TaobaoTbkSpreadGetRequest{
    return &TaobaoTbkSpreadGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkSpreadGetRequest) GetApiMethodName() string {
    return "taobao.tbk.spread.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkSpreadGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Requests Setter
// 请求列表，内部包含多个url
func (r *TaobaoTbkSpreadGetRequest) SetRequests(requests []TbkSpreadRequest) error {
    r.requests = requests
    r.Set("requests", requests)
    return nil
}

// Requests Getter
func (r TaobaoTbkSpreadGetRequest) GetRequests() []TbkSpreadRequest {
    return r.requests
}
