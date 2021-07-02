package seaking

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSeakingServicepackAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.servicepack"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSeakingServicepackAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
