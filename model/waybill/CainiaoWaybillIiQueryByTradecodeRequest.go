package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过订单号查询电子面单通接口 APIRequest
cainiao.waybill.ii.query.by.tradecode

通过订单号查看面单的信息
*/
type CainiaoWaybillIiQueryByTradecodeRequest struct {
    model.Params

    // 订单号列表
    paramList   []WaybillDetailQueryByBizSubCodeRequest 

}

func NewCainiaoWaybillIiQueryByTradecodeRequest() *CainiaoWaybillIiQueryByTradecodeRequest{
    return &CainiaoWaybillIiQueryByTradecodeRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoWaybillIiQueryByTradecodeRequest) GetApiMethodName() string {
    return "cainiao.waybill.ii.query.by.tradecode"
}

func (r CainiaoWaybillIiQueryByTradecodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoWaybillIiQueryByTradecodeRequest) SetParamList(paramList []WaybillDetailQueryByBizSubCodeRequest) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

func (r CainiaoWaybillIiQueryByTradecodeRequest) GetParamList() []WaybillDetailQueryByBizSubCodeRequest {
    return r.paramList
}

