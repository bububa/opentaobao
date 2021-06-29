package fivee

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商信息 API请求
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

// 初始化TaobaoFiveeCompanyGetRequest对象
func NewTaobaoFiveeCompanyGetRequest() *TaobaoFiveeCompanyGetRequest{
    return &TaobaoFiveeCompanyGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFiveeCompanyGetRequest) GetApiMethodName() string {
    return "taobao.fivee.company.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFiveeCompanyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBucode Setter
// bu身份标识
func (r *TaobaoFiveeCompanyGetRequest) SetParamBucode(paramBucode string) error {
    r.paramBucode = paramBucode
    r.Set("param_bucode", paramBucode)
    return nil
}

// ParamBucode Getter
func (r TaobaoFiveeCompanyGetRequest) GetParamBucode() string {
    return r.paramBucode
}
// ParamUniqueCode Setter
// 统一社会信息用代码
func (r *TaobaoFiveeCompanyGetRequest) SetParamUniqueCode(paramUniqueCode string) error {
    r.paramUniqueCode = paramUniqueCode
    r.Set("param_unique_code", paramUniqueCode)
    return nil
}

// ParamUniqueCode Getter
func (r TaobaoFiveeCompanyGetRequest) GetParamUniqueCode() string {
    return r.paramUniqueCode
}
