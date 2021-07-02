package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest 用户实名认证 API请求
// taobao.jym.member.realname.verify.withoutuid
//
// 开放用户实名认证接口使用
type TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest struct {
	model.Params
	// 加密名字串
	_encryptName string
	// 加密身份证串
	_encryptIdNo string
}

// NewTaobaoJymMemberRealnameVerifyWithoutuidRequest 初始化TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest对象
func NewTaobaoJymMemberRealnameVerifyWithoutuidRequest() *TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest {
	return &TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest) GetApiMethodName() string {
	return "taobao.jym.member.realname.verify.withoutuid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is EncryptName Setter
// 加密名字串
func (r *TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest) SetEncryptName(_encryptName string) error {
	r._encryptName = _encryptName
	r.Set("encrypt_name", _encryptName)
	return nil
}

// Get EncryptName Getter
func (r TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest) GetEncryptName() string {
	return r._encryptName
}

// Set is EncryptIdNo Setter
// 加密身份证串
func (r *TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest) SetEncryptIdNo(_encryptIdNo string) error {
	r._encryptIdNo = _encryptIdNo
	r.Set("encrypt_id_no", _encryptIdNo)
	return nil
}

// Get EncryptIdNo Getter
func (r TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest) GetEncryptIdNo() string {
	return r._encryptIdNo
}
