package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据outerCode查询商品 API请求
taobao.scitem.outercode.get

根据outerCode查询商品
*/
type TaobaoScitemOutercodeGetAPIRequest struct {
    model.Params
    // 商品编码
    _outerCode   string
}

// 初始化TaobaoScitemOutercodeGetAPIRequest对象
func NewTaobaoScitemOutercodeGetRequest() *TaobaoScitemOutercodeGetAPIRequest{
    return &TaobaoScitemOutercodeGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoScitemOutercodeGetAPIRequest) GetApiMethodName() string {
    return "taobao.scitem.outercode.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoScitemOutercodeGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterCode Setter
// 商品编码
func (r *TaobaoScitemOutercodeGetAPIRequest) SetOuterCode(_outerCode string) error {
    r._outerCode = _outerCode
    r.Set("outer_code", _outerCode)
    return nil
}

// OuterCode Getter
func (r TaobaoScitemOutercodeGetAPIRequest) GetOuterCode() string {
    return r._outerCode
}
