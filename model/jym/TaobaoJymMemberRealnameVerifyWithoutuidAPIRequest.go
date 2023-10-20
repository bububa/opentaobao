package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojymmemberrealnameverifywithoutuidAPIRequest 用户实名认证 API请求
// taobao.jym.member.realname.verify.withoutuid
//
// 开放用户实名认证接口使用
type TaobaojymmemberrealnameverifywithoutuidAPIRequest struct {
	model.Params
	// 加密名字串
	_encryptName string
	// 加密身份证串
	_encryptIdNo string
}

// NewTaobaojymmemberrealnameverifywithoutuidRequest 初始化TaobaojymmemberrealnameverifywithoutuidAPIRequest对象
func NewTaobaojymmemberrealnameverifywithoutuidRequest() *TaobaojymmemberrealnameverifywithoutuidAPIRequest {
	return &TaobaojymmemberrealnameverifywithoutuidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojymmemberrealnameverifywithoutuidAPIRequest) GetApiMethodName() string {
	return "taobao.jym.member.realname.verify.withoutuid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojymmemberrealnameverifywithoutuidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojymmemberrealnameverifywithoutuidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEncryptName is EncryptName Setter
// 加密名字串
func (r *TaobaojymmemberrealnameverifywithoutuidAPIRequest) SetEncryptName(_encryptName string) error {
	r._encryptName = _encryptName
	r.Set("encrypt_name", _encryptName)
	return nil
}

// GetEncryptName EncryptName Getter
func (r TaobaojymmemberrealnameverifywithoutuidAPIRequest) GetEncryptName() string {
	return r._encryptName
}

// SetEncryptIdNo is EncryptIdNo Setter
// 加密身份证串
func (r *TaobaojymmemberrealnameverifywithoutuidAPIRequest) SetEncryptIdNo(_encryptIdNo string) error {
	r._encryptIdNo = _encryptIdNo
	r.Set("encrypt_id_no", _encryptIdNo)
	return nil
}

// GetEncryptIdNo EncryptIdNo Getter
func (r TaobaojymmemberrealnameverifywithoutuidAPIRequest) GetEncryptIdNo() string {
	return r._encryptIdNo
}
