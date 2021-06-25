package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询解析淘口令 APIRequest
alibaba.baichuan.taopassword.query

查询，解析淘口令
*/
type AlibabaBaichuanTaopasswordQueryRequest struct {
    model.Params

    // 淘口令
    passwordContent   string 

}

func NewAlibabaBaichuanTaopasswordQueryRequest() *AlibabaBaichuanTaopasswordQueryRequest{
    return &AlibabaBaichuanTaopasswordQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaBaichuanTaopasswordQueryRequest) GetApiMethodName() string {
    return "alibaba.baichuan.taopassword.query"
}

func (r AlibabaBaichuanTaopasswordQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaBaichuanTaopasswordQueryRequest) SetPasswordContent(passwordContent string) error {
    r.passwordContent = passwordContent
    r.Set("password_content", passwordContent)
    return nil
}

func (r AlibabaBaichuanTaopasswordQueryRequest) GetPasswordContent() string {
    return r.passwordContent
}

