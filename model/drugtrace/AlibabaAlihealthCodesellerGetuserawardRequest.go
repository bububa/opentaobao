package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机扫码查询领奖状态 APIRequest
alibaba.alihealth.codeseller.getuseraward

贩卖机扫码查询领奖状态
*/
type AlibabaAlihealthCodesellerGetuserawardRequest struct {
    model.Params

    // 追溯码
    code   string 

}

func NewAlibabaAlihealthCodesellerGetuserawardRequest() *AlibabaAlihealthCodesellerGetuserawardRequest{
    return &AlibabaAlihealthCodesellerGetuserawardRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthCodesellerGetuserawardRequest) GetApiMethodName() string {
    return "alibaba.alihealth.codeseller.getuseraward"
}

func (r AlibabaAlihealthCodesellerGetuserawardRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthCodesellerGetuserawardRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthCodesellerGetuserawardRequest) GetCode() string {
    return r.code
}

