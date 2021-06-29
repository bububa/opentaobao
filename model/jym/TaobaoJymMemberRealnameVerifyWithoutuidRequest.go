package jym

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户实名认证 APIRequest
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

func NewTaobaoJymMemberRealnameVerifyWithoutuidRequest() *TaobaoJymMemberRealnameVerifyWithoutuidRequest{
    return &TaobaoJymMemberRealnameVerifyWithoutuidRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJymMemberRealnameVerifyWithoutuidRequest) GetApiMethodName() string {
    return "taobao.jym.member.realname.verify.withoutuid"
}

func (r TaobaoJymMemberRealnameVerifyWithoutuidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJymMemberRealnameVerifyWithoutuidRequest) SetEncryptName(encryptName string) error {
    r.encryptName = encryptName
    r.Set("encrypt_name", encryptName)
    return nil
}

func (r TaobaoJymMemberRealnameVerifyWithoutuidRequest) GetEncryptName() string {
    return r.encryptName
}

func (r *TaobaoJymMemberRealnameVerifyWithoutuidRequest) SetEncryptIdNo(encryptIdNo string) error {
    r.encryptIdNo = encryptIdNo
    r.Set("encrypt_id_no", encryptIdNo)
    return nil
}

func (r TaobaoJymMemberRealnameVerifyWithoutuidRequest) GetEncryptIdNo() string {
    return r.encryptIdNo
}

