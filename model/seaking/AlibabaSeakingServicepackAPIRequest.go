package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaseakingservicepackAPIRequest 获取海王用户权限包 API请求
// alibaba.seaking.servicepack
//
// 获取海王用户权限包
type AlibabaseakingservicepackAPIRequest struct {
	model.Params
	// 验证类型
	_identifyType string
	// 验证类型下的唯一id
	_identifier string
}

// NewAlibabaseakingservicepackRequest 初始化AlibabaseakingservicepackAPIRequest对象
func NewAlibabaseakingservicepackRequest() *AlibabaseakingservicepackAPIRequest {
	return &AlibabaseakingservicepackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaseakingservicepackAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.servicepack"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaseakingservicepackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaseakingservicepackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdentifyType is IdentifyType Setter
// 验证类型
func (r *AlibabaseakingservicepackAPIRequest) SetIdentifyType(_identifyType string) error {
	r._identifyType = _identifyType
	r.Set("identify_type", _identifyType)
	return nil
}

// GetIdentifyType IdentifyType Getter
func (r AlibabaseakingservicepackAPIRequest) GetIdentifyType() string {
	return r._identifyType
}

// SetIdentifier is Identifier Setter
// 验证类型下的唯一id
func (r *AlibabaseakingservicepackAPIRequest) SetIdentifier(_identifier string) error {
	r._identifier = _identifier
	r.Set("identifier", _identifier)
	return nil
}

// GetIdentifier Identifier Getter
func (r AlibabaseakingservicepackAPIRequest) GetIdentifier() string {
	return r._identifier
}
