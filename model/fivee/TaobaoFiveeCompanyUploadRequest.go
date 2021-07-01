package fivee

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传商信息接口 API请求
taobao.fivee.company.upload

资质共享平台上传资质证照
*/
type TaobaoFiveeCompanyUploadAPIRequest struct {
    model.Params
    // bu身份标识
    _paramBucode   string
    // 商家证照信息
    _paramCompany   *Company
}

// 初始化TaobaoFiveeCompanyUploadAPIRequest对象
func NewTaobaoFiveeCompanyUploadRequest() *TaobaoFiveeCompanyUploadAPIRequest{
    return &TaobaoFiveeCompanyUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFiveeCompanyUploadAPIRequest) GetApiMethodName() string {
    return "taobao.fivee.company.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFiveeCompanyUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBucode Setter
// bu身份标识
func (r *TaobaoFiveeCompanyUploadAPIRequest) SetParamBucode(_paramBucode string) error {
    r._paramBucode = _paramBucode
    r.Set("param_bucode", _paramBucode)
    return nil
}

// ParamBucode Getter
func (r TaobaoFiveeCompanyUploadAPIRequest) GetParamBucode() string {
    return r._paramBucode
}
// ParamCompany Setter
// 商家证照信息
func (r *TaobaoFiveeCompanyUploadAPIRequest) SetParamCompany(_paramCompany *Company) error {
    r._paramCompany = _paramCompany
    r.Set("param_company", _paramCompany)
    return nil
}

// ParamCompany Getter
func (r TaobaoFiveeCompanyUploadAPIRequest) GetParamCompany() *Company {
    return r._paramCompany
}
