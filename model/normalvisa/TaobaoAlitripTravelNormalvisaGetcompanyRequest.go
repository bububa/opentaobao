package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取物流公司信息 APIRequest
taobao.alitrip.travel.normalvisa.getcompany

获取物流公司信息
*/
type TaobaoAlitripTravelNormalvisaGetcompanyRequest struct {
    model.Params

    // true：取5个重要的物流公司 false：取所有的物流公司
    param0   bool 

}

func NewTaobaoAlitripTravelNormalvisaGetcompanyRequest() *TaobaoAlitripTravelNormalvisaGetcompanyRequest{
    return &TaobaoAlitripTravelNormalvisaGetcompanyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTravelNormalvisaGetcompanyRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.normalvisa.getcompany"
}

func (r TaobaoAlitripTravelNormalvisaGetcompanyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTravelNormalvisaGetcompanyRequest) SetParam0(param0 bool) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoAlitripTravelNormalvisaGetcompanyRequest) GetParam0() bool {
    return r.param0
}

