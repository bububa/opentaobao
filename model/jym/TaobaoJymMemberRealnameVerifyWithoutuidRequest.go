package jym

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户实名认证 API请求
taobao.jym.member.realname.verify.withoutuid

开放用户实名认证接口使用
*/
type TaobaoJymMemberRealnameVerifyWithoutuidRequest struct {
    model.Params
    // 加密名字串
    encryptName   string
    // 加密身份证串
    encryptIdNo   string
}

// 初始化TaobaoJymMemberRealnameVerifyWithoutuidRequest对象
func NewTaobaoJymMemberRealnameVerifyWithoutuidRequest() *TaobaoJymMemberRealnameVerifyWithoutuidRequest{
    return &TaobaoJymMemberRealnameVerifyWithoutuidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJymMemberRealnameVerifyWithoutuidRequest) GetApiMethodName() string {
    return "taobao.jym.member.realname.verify.withoutuid"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJymMemberRealnameVerifyWithoutuidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EncryptName Setter
// 加密名字串
func (r *TaobaoJymMemberRealnameVerifyWithoutuidRequest) SetEncryptName(encryptName string) error {
    r.encryptName = encryptName
    r.Set("encrypt_name", encryptName)
    return nil
}

// EncryptName Getter
func (r TaobaoJymMemberRealnameVerifyWithoutuidRequest) GetEncryptName() string {
    return r.encryptName
}
// EncryptIdNo Setter
// 加密身份证串
func (r *TaobaoJymMemberRealnameVerifyWithoutuidRequest) SetEncryptIdNo(encryptIdNo string) error {
    r.encryptIdNo = encryptIdNo
    r.Set("encrypt_id_no", encryptIdNo)
    return nil
}

// EncryptIdNo Getter
func (r TaobaoJymMemberRealnameVerifyWithoutuidRequest) GetEncryptIdNo() string {
    return r.encryptIdNo
}
