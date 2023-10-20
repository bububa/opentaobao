package jym

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest) Reset() {
	r._encryptName = ""
	r._encryptIdNo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest) GetApiMethodName() string {
	return "taobao.jym.member.realname.verify.withoutuid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEncryptName is EncryptName Setter
// 加密名字串
func (r *TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest) SetEncryptName(_encryptName string) error {
	r._encryptName = _encryptName
	r.Set("encrypt_name", _encryptName)
	return nil
}

// GetEncryptName EncryptName Getter
func (r TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest) GetEncryptName() string {
	return r._encryptName
}

// SetEncryptIdNo is EncryptIdNo Setter
// 加密身份证串
func (r *TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest) SetEncryptIdNo(_encryptIdNo string) error {
	r._encryptIdNo = _encryptIdNo
	r.Set("encrypt_id_no", _encryptIdNo)
	return nil
}

// GetEncryptIdNo EncryptIdNo Getter
func (r TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest) GetEncryptIdNo() string {
	return r._encryptIdNo
}

var poolTaobaoJymMemberRealnameVerifyWithoutuidAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJymMemberRealnameVerifyWithoutuidRequest()
	},
}

// GetTaobaoJymMemberRealnameVerifyWithoutuidRequest 从 sync.Pool 获取 TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest
func GetTaobaoJymMemberRealnameVerifyWithoutuidAPIRequest() *TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest {
	return poolTaobaoJymMemberRealnameVerifyWithoutuidAPIRequest.Get().(*TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest)
}

// ReleaseTaobaoJymMemberRealnameVerifyWithoutuidAPIRequest 将 TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest 放入 sync.Pool
func ReleaseTaobaoJymMemberRealnameVerifyWithoutuidAPIRequest(v *TaobaoJymMemberRealnameVerifyWithoutuidAPIRequest) {
	v.Reset()
	poolTaobaoJymMemberRealnameVerifyWithoutuidAPIRequest.Put(v)
}
