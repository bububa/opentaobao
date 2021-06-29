package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询解析淘口令 API请求
alibaba.baichuan.taopassword.query

查询，解析淘口令
*/
type AlibabaBaichuanTaopasswordQueryRequest struct {
    model.Params
    // 淘口令
    passwordContent   string
}

// 初始化AlibabaBaichuanTaopasswordQueryRequest对象
func NewAlibabaBaichuanTaopasswordQueryRequest() *AlibabaBaichuanTaopasswordQueryRequest{
    return &AlibabaBaichuanTaopasswordQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanTaopasswordQueryRequest) GetApiMethodName() string {
    return "alibaba.baichuan.taopassword.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanTaopasswordQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PasswordContent Setter
// 淘口令
func (r *AlibabaBaichuanTaopasswordQueryRequest) SetPasswordContent(passwordContent string) error {
    r.passwordContent = passwordContent
    r.Set("password_content", passwordContent)
    return nil
}

// PasswordContent Getter
func (r AlibabaBaichuanTaopasswordQueryRequest) GetPasswordContent() string {
    return r.passwordContent
}
