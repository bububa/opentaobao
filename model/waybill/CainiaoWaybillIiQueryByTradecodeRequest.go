package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过订单号查询电子面单通接口 API请求
cainiao.waybill.ii.query.by.tradecode

通过订单号查看面单的信息
*/
type CainiaoWaybillIiQueryByTradecodeRequest struct {
    model.Params
    // 订单号列表
    _paramList   []WaybillDetailQueryByBizSubCodeRequest
}

// 初始化CainiaoWaybillIiQueryByTradecodeRequest对象
func NewCainiaoWaybillIiQueryByTradecodeRequest() *CainiaoWaybillIiQueryByTradecodeRequest{
    return &CainiaoWaybillIiQueryByTradecodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiQueryByTradecodeRequest) GetApiMethodName() string {
    return "cainiao.waybill.ii.query.by.tradecode"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiQueryByTradecodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamList Setter
// 订单号列表
func (r *CainiaoWaybillIiQueryByTradecodeRequest) SetParamList(_paramList []WaybillDetailQueryByBizSubCodeRequest) error {
    r._paramList = _paramList
    r.Set("param_list", _paramList)
    return nil
}

// ParamList Getter
func (r CainiaoWaybillIiQueryByTradecodeRequest) GetParamList() []WaybillDetailQueryByBizSubCodeRequest {
    return r._paramList
}
