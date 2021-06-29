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
type TaobaoScitemOutercodeGetRequest struct {
    model.Params
    // 商品编码
    _outerCode   string
}

// 初始化TaobaoScitemOutercodeGetRequest对象
func NewTaobaoScitemOutercodeGetRequest() *TaobaoScitemOutercodeGetRequest{
    return &TaobaoScitemOutercodeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoScitemOutercodeGetRequest) GetApiMethodName() string {
    return "taobao.scitem.outercode.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoScitemOutercodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterCode Setter
// 商品编码
func (r *TaobaoScitemOutercodeGetRequest) SetOuterCode(_outerCode string) error {
    r._outerCode = _outerCode
    r.Set("outer_code", _outerCode)
    return nil
}

// OuterCode Getter
func (r TaobaoScitemOutercodeGetRequest) GetOuterCode() string {
    return r._outerCode
}
