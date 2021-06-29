package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改外部成本中心 API请求
alitrip.btrip.cost.center.modify

修改外部成本中心，设置成员，设置支付宝账号，设置名称，编号等
*/
type AlitripBtripCostCenterModifyRequest struct {
    model.Params
    // 请求对象
    rq   *OpenCostCenterModifyRq
}

// 初始化AlitripBtripCostCenterModifyRequest对象
func NewAlitripBtripCostCenterModifyRequest() *AlitripBtripCostCenterModifyRequest{
    return &AlitripBtripCostCenterModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterModifyRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.modify"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripCostCenterModifyRequest) SetRq(rq *OpenCostCenterModifyRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCostCenterModifyRequest) GetRq() *OpenCostCenterModifyRq {
    return r.rq
}
