package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据outerCode查询商品 APIRequest
taobao.scitem.outercode.get

根据outerCode查询商品
*/
type TaobaoScitemOutercodeGetRequest struct {
    model.Params

    // 商品编码
    outerCode   string 

}

func NewTaobaoScitemOutercodeGetRequest() *TaobaoScitemOutercodeGetRequest{
    return &TaobaoScitemOutercodeGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoScitemOutercodeGetRequest) GetApiMethodName() string {
    return "taobao.scitem.outercode.get"
}

func (r TaobaoScitemOutercodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoScitemOutercodeGetRequest) SetOuterCode(outerCode string) error {
    r.outerCode = outerCode
    r.Set("outer_code", outerCode)
    return nil
}

func (r TaobaoScitemOutercodeGetRequest) GetOuterCode() string {
    return r.outerCode
}

