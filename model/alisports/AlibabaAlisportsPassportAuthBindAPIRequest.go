package alisports

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportAuthBindAPIRequest 授权绑定关系接口 API请求
// alibaba.alisports.passport.auth.bind
//
// 授权回调绑定关系接口，建立阿里体育openId和三方openId的绑定关系
type AlibabaAlisportsPassportAuthBindAPIRequest struct {
	model.Params
	// 阿里体育业务KEY
	_alispAppKey string
	// 阿里体育openId
	_openId string
	// 合作方openId
	_thirdOpenId string
}

// NewAlibabaAlisportsPassportAuthBindRequest 初始化AlibabaAlisportsPassportAuthBindAPIRequest对象
func NewAlibabaAlisportsPassportAuthBindRequest() *AlibabaAlisportsPassportAuthBindAPIRequest {
	return &AlibabaAlisportsPassportAuthBindAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlisportsPassportAuthBindAPIRequest) Reset() {
	r._alispAppKey = ""
	r._openId = ""
	r._thirdOpenId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportAuthBindAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.passport.auth.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportAuthBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlisportsPassportAuthBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlispAppKey is AlispAppKey Setter
// 阿里体育业务KEY
func (r *AlibabaAlisportsPassportAuthBindAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaAlisportsPassportAuthBindAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetOpenId is OpenId Setter
// 阿里体育openId
func (r *AlibabaAlisportsPassportAuthBindAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r AlibabaAlisportsPassportAuthBindAPIRequest) GetOpenId() string {
	return r._openId
}

// SetThirdOpenId is ThirdOpenId Setter
// 合作方openId
func (r *AlibabaAlisportsPassportAuthBindAPIRequest) SetThirdOpenId(_thirdOpenId string) error {
	r._thirdOpenId = _thirdOpenId
	r.Set("third_open_id", _thirdOpenId)
	return nil
}

// GetThirdOpenId ThirdOpenId Getter
func (r AlibabaAlisportsPassportAuthBindAPIRequest) GetThirdOpenId() string {
	return r._thirdOpenId
}

var poolAlibabaAlisportsPassportAuthBindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlisportsPassportAuthBindRequest()
	},
}

// GetAlibabaAlisportsPassportAuthBindRequest 从 sync.Pool 获取 AlibabaAlisportsPassportAuthBindAPIRequest
func GetAlibabaAlisportsPassportAuthBindAPIRequest() *AlibabaAlisportsPassportAuthBindAPIRequest {
	return poolAlibabaAlisportsPassportAuthBindAPIRequest.Get().(*AlibabaAlisportsPassportAuthBindAPIRequest)
}

// ReleaseAlibabaAlisportsPassportAuthBindAPIRequest 将 AlibabaAlisportsPassportAuthBindAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlisportsPassportAuthBindAPIRequest(v *AlibabaAlisportsPassportAuthBindAPIRequest) {
	v.Reset()
	poolAlibabaAlisportsPassportAuthBindAPIRequest.Put(v)
}
