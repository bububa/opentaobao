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
type AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest struct {
    model.Params
    // WorkBenchContext
    _param0   *WorkBenchContext
}

// 初始化AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest对象
func NewAlibabaCampusCoreCompanycampusGetcombycamidRequest() *AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest{
    return &AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.core.companycampus.getcombycamid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// WorkBenchContext
func (r *AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusCoreCompanycampusGetcombycamidAPIRequest) GetParam0() *WorkBenchContext {
    return r._param0
}
