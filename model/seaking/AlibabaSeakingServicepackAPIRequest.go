package seaking

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingServicepackAPIRequest 获取海王用户权限包 API请求
// alibaba.seaking.servicepack
//
// 获取海王用户权限包
type AlibabaSeakingServicepackAPIRequest struct {
	model.Params
	// 验证类型
	_identifyType string
	// 验证类型下的唯一id
	_identifier string
}

// NewAlibabaSeakingServicepackRequest 初始化AlibabaSeakingServicepackAPIRequest对象
func NewAlibabaSeakingServicepackRequest() *AlibabaSeakingServicepackAPIRequest {
	return &AlibabaSeakingServicepackAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSeakingServicepackAPIRequest) Reset() {
	r._identifyType = ""
	r._identifier = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSeakingServicepackAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.servicepack"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSeakingServicepackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSeakingServicepackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdentifyType is IdentifyType Setter
// 验证类型
func (r *AlibabaSeakingServicepackAPIRequest) SetIdentifyType(_identifyType string) error {
	r._identifyType = _identifyType
	r.Set("identify_type", _identifyType)
	return nil
}

// GetIdentifyType IdentifyType Getter
func (r AlibabaSeakingServicepackAPIRequest) GetIdentifyType() string {
	return r._identifyType
}

// SetIdentifier is Identifier Setter
// 验证类型下的唯一id
func (r *AlibabaSeakingServicepackAPIRequest) SetIdentifier(_identifier string) error {
	r._identifier = _identifier
	r.Set("identifier", _identifier)
	return nil
}

// GetIdentifier Identifier Getter
func (r AlibabaSeakingServicepackAPIRequest) GetIdentifier() string {
	return r._identifier
}

var poolAlibabaSeakingServicepackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSeakingServicepackRequest()
	},
}

// GetAlibabaSeakingServicepackRequest 从 sync.Pool 获取 AlibabaSeakingServicepackAPIRequest
func GetAlibabaSeakingServicepackAPIRequest() *AlibabaSeakingServicepackAPIRequest {
	return poolAlibabaSeakingServicepackAPIRequest.Get().(*AlibabaSeakingServicepackAPIRequest)
}

// ReleaseAlibabaSeakingServicepackAPIRequest 将 AlibabaSeakingServicepackAPIRequest 放入 sync.Pool
func ReleaseAlibabaSeakingServicepackAPIRequest(v *AlibabaSeakingServicepackAPIRequest) {
	v.Reset()
	poolAlibabaSeakingServicepackAPIRequest.Put(v)
}
