package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmMarketingEncryptAPIRequest 加密 API请求
// alibaba.alsc.crm.marketing.encrypt
//
// 加密
type AlibabaAlscCrmMarketingEncryptAPIRequest struct {
	model.Params
	// 参数
	_param string
}

// NewAlibabaAlscCrmMarketingEncryptRequest 初始化AlibabaAlscCrmMarketingEncryptAPIRequest对象
func NewAlibabaAlscCrmMarketingEncryptRequest() *AlibabaAlscCrmMarketingEncryptAPIRequest {
	return &AlibabaAlscCrmMarketingEncryptAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmMarketingEncryptAPIRequest) Reset() {
	r._param = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmMarketingEncryptAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.marketing.encrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmMarketingEncryptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmMarketingEncryptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 参数
func (r *AlibabaAlscCrmMarketingEncryptAPIRequest) SetParam(_param string) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaAlscCrmMarketingEncryptAPIRequest) GetParam() string {
	return r._param
}

var poolAlibabaAlscCrmMarketingEncryptAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmMarketingEncryptRequest()
	},
}

// GetAlibabaAlscCrmMarketingEncryptRequest 从 sync.Pool 获取 AlibabaAlscCrmMarketingEncryptAPIRequest
func GetAlibabaAlscCrmMarketingEncryptAPIRequest() *AlibabaAlscCrmMarketingEncryptAPIRequest {
	return poolAlibabaAlscCrmMarketingEncryptAPIRequest.Get().(*AlibabaAlscCrmMarketingEncryptAPIRequest)
}

// ReleaseAlibabaAlscCrmMarketingEncryptAPIRequest 将 AlibabaAlscCrmMarketingEncryptAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmMarketingEncryptAPIRequest(v *AlibabaAlscCrmMarketingEncryptAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmMarketingEncryptAPIRequest.Put(v)
}
