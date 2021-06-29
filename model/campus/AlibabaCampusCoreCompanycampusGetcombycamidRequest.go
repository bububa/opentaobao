package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据园区ID获取运营公司信息 APIRequest
alibaba.campus.core.companycampus.getcombycamid

根据园区ID获取运营公司信息
*/
type AlibabaCampusCoreCompanycampusGetcombycamidRequest struct {
    model.Params

    // WorkBenchContext
    param0   *WorkBenchContext 

}

func NewAlibabaCampusCoreCompanycampusGetcombycamidRequest() *AlibabaCampusCoreCompanycampusGetcombycamidRequest{
    return &AlibabaCampusCoreCompanycampusGetcombycamidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusCoreCompanycampusGetcombycamidRequest) GetApiMethodName() string {
    return "alibaba.campus.core.companycampus.getcombycamid"
}

func (r AlibabaCampusCoreCompanycampusGetcombycamidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusCoreCompanycampusGetcombycamidRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaCampusCoreCompanycampusGetcombycamidRequest) GetParam0() *WorkBenchContext {
    return r.param0
}

