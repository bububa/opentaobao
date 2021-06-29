package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取物流公司信息 API请求
taobao.alitrip.travel.normalvisa.getcompany

获取物流公司信息
*/
type TaobaoAlitripTravelNormalvisaGetcompanyRequest struct {
    model.Params
    // true：取5个重要的物流公司 false：取所有的物流公司
    param0   bool
}

// 初始化TaobaoAlitripTravelNormalvisaGetcompanyRequest对象
func NewTaobaoAlitripTravelNormalvisaGetcompanyRequest() *TaobaoAlitripTravelNormalvisaGetcompanyRequest{
    return &TaobaoAlitripTravelNormalvisaGetcompanyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelNormalvisaGetcompanyRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.normalvisa.getcompany"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelNormalvisaGetcompanyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// true：取5个重要的物流公司 false：取所有的物流公司
func (r *TaobaoAlitripTravelNormalvisaGetcompanyRequest) SetParam0(param0 bool) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r TaobaoAlitripTravelNormalvisaGetcompanyRequest) GetParam0() bool {
    return r.param0
}
