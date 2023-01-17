package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingAuthmachineapiAPIRequest 机翻Api授权 API请求
// alibaba.seaking.authmachineapi
//
// 机翻Api授权
type AlibabaSeakingAuthmachineapiAPIRequest struct {
	model.Params
	// erp名称
	_identifyType string
	// erp用户id
	_identifier string
	// 店铺所在平台
	_subIdentifyType string
	// 店铺id(ae为cn开头的店铺id, lazada为邮箱)
	_subIdentifier string
}

// NewAlibabaSeakingAuthmachineapiRequest 初始化AlibabaSeakingAuthmachineapiAPIRequest对象
func NewAlibabaSeakingAuthmachineapiRequest() *AlibabaSeakingAuthmachineapiAPIRequest {
	return &AlibabaSeakingAuthmachineapiAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSeakingAuthmachineapiAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.authmachineapi"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSeakingAuthmachineapiAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSeakingAuthmachineapiAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdentifyType is IdentifyType Setter
// erp名称
func (r *AlibabaSeakingAuthmachineapiAPIRequest) SetIdentifyType(_identifyType string) error {
	r._identifyType = _identifyType
	r.Set("identify_type", _identifyType)
	return nil
}

// GetIdentifyType IdentifyType Getter
func (r AlibabaSeakingAuthmachineapiAPIRequest) GetIdentifyType() string {
	return r._identifyType
}

// SetIdentifier is Identifier Setter
// erp用户id
func (r *AlibabaSeakingAuthmachineapiAPIRequest) SetIdentifier(_identifier string) error {
	r._identifier = _identifier
	r.Set("identifier", _identifier)
	return nil
}

// GetIdentifier Identifier Getter
func (r AlibabaSeakingAuthmachineapiAPIRequest) GetIdentifier() string {
	return r._identifier
}

// SetSubIdentifyType is SubIdentifyType Setter
// 店铺所在平台
func (r *AlibabaSeakingAuthmachineapiAPIRequest) SetSubIdentifyType(_subIdentifyType string) error {
	r._subIdentifyType = _subIdentifyType
	r.Set("sub_identify_type", _subIdentifyType)
	return nil
}

// GetSubIdentifyType SubIdentifyType Getter
func (r AlibabaSeakingAuthmachineapiAPIRequest) GetSubIdentifyType() string {
	return r._subIdentifyType
}

// SetSubIdentifier is SubIdentifier Setter
// 店铺id(ae为cn开头的店铺id, lazada为邮箱)
func (r *AlibabaSeakingAuthmachineapiAPIRequest) SetSubIdentifier(_subIdentifier string) error {
	r._subIdentifier = _subIdentifier
	r.Set("sub_identifier", _subIdentifier)
	return nil
}

// GetSubIdentifier SubIdentifier Getter
func (r AlibabaSeakingAuthmachineapiAPIRequest) GetSubIdentifier() string {
	return r._subIdentifier
}
