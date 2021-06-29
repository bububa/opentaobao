package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据园区ID获取运营公司信息 API请求
alibaba.campus.core.companycampus.getcombycamid

根据园区ID获取运营公司信息
*/
type AlibabaCampusCoreCompanycampusGetcombycamidRequest struct {
    model.Params
    // WorkBenchContext
    _param0   *WorkBenchContext
}

// 初始化AlibabaCampusCoreCompanycampusGetcombycamidRequest对象
func NewAlibabaCampusCoreCompanycampusGetcombycamidRequest() *AlibabaCampusCoreCompanycampusGetcombycamidRequest{
    return &AlibabaCampusCoreCompanycampusGetcombycamidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusCoreCompanycampusGetcombycamidRequest) GetApiMethodName() string {
    return "alibaba.campus.core.companycampus.getcombycamid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusCoreCompanycampusGetcombycamidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// WorkBenchContext
func (r *AlibabaCampusCoreCompanycampusGetcombycamidRequest) SetParam0(_param0 *WorkBenchContext) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusCoreCompanycampusGetcombycamidRequest) GetParam0() *WorkBenchContext {
    return r._param0
}
