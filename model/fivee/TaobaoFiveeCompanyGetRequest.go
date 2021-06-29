package fivee

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商信息 APIRequest
taobao.fivee.company.get

资质共享平台查询商信息
*/
type TaobaoFiveeCompanyGetRequest struct {
    model.Params

    // bu身份标识
    paramBucode   string 

    // 统一社会信息用代码
    paramUniqueCode   string 

}

func NewTaobaoFiveeCompanyGetRequest() *TaobaoFiveeCompanyGetRequest{
    return &TaobaoFiveeCompanyGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFiveeCompanyGetRequest) GetApiMethodName() string {
    return "taobao.fivee.company.get"
}

func (r TaobaoFiveeCompanyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFiveeCompanyGetRequest) SetParamBucode(paramBucode string) error {
    r.paramBucode = paramBucode
    r.Set("param_bucode", paramBucode)
    return nil
}

func (r TaobaoFiveeCompanyGetRequest) GetParamBucode() string {
    return r.paramBucode
}

func (r *TaobaoFiveeCompanyGetRequest) SetParamUniqueCode(paramUniqueCode string) error {
    r.paramUniqueCode = paramUniqueCode
    r.Set("param_unique_code", paramUniqueCode)
    return nil
}

func (r TaobaoFiveeCompanyGetRequest) GetParamUniqueCode() string {
    return r.paramUniqueCode
}

