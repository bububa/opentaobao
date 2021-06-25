package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过面单号查询面单信息 APIRequest
cainiao.waybill.ii.query.by.waybillcode

通过面单号查看面单号的当前状态，如签收、发货、失效等。
*/
type CainiaoWaybillIiQueryByWaybillcodeRequest struct {
    model.Params

    // 系统自动生成
    paramList   []WaybillDetailQueryByWaybillCodeRequest 

}

func NewCainiaoWaybillIiQueryByWaybillcodeRequest() *CainiaoWaybillIiQueryByWaybillcodeRequest{
    return &CainiaoWaybillIiQueryByWaybillcodeRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoWaybillIiQueryByWaybillcodeRequest) GetApiMethodName() string {
    return "cainiao.waybill.ii.query.by.waybillcode"
}

func (r CainiaoWaybillIiQueryByWaybillcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoWaybillIiQueryByWaybillcodeRequest) SetParamList(paramList []WaybillDetailQueryByWaybillCodeRequest) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

func (r CainiaoWaybillIiQueryByWaybillcodeRequest) GetParamList() []WaybillDetailQueryByWaybillCodeRequest {
    return r.paramList
}

