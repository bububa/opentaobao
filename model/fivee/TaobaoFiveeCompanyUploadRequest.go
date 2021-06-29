package fivee

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传商信息接口 APIRequest
taobao.fivee.company.upload

资质共享平台上传资质证照
*/
type TaobaoFiveeCompanyUploadRequest struct {
    model.Params

    // bu身份标识
    paramBucode   string 

    // 商家证照信息
    paramCompany   *Company 

}

func NewTaobaoFiveeCompanyUploadRequest() *TaobaoFiveeCompanyUploadRequest{
    return &TaobaoFiveeCompanyUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFiveeCompanyUploadRequest) GetApiMethodName() string {
    return "taobao.fivee.company.upload"
}

func (r TaobaoFiveeCompanyUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFiveeCompanyUploadRequest) SetParamBucode(paramBucode string) error {
    r.paramBucode = paramBucode
    r.Set("param_bucode", paramBucode)
    return nil
}

func (r TaobaoFiveeCompanyUploadRequest) GetParamBucode() string {
    return r.paramBucode
}

func (r *TaobaoFiveeCompanyUploadRequest) SetParamCompany(paramCompany *Company) error {
    r.paramCompany = paramCompany
    r.Set("param_company", paramCompany)
    return nil
}

func (r TaobaoFiveeCompanyUploadRequest) GetParamCompany() *Company {
    return r.paramCompany
}

