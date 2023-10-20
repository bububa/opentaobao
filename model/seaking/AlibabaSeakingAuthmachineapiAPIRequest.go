package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaseakingauthmachineapiAPIRequest 机翻Api授权 API请求
// alibaba.seaking.authmachineapi
//
// 机翻Api授权
type AlibabaseakingauthmachineapiAPIRequest struct {
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

// NewAlibabaseakingauthmachineapiRequest 初始化AlibabaseakingauthmachineapiAPIRequest对象
func NewAlibabaseakingauthmachineapiRequest() *AlibabaseakingauthmachineapiAPIRequest {
	return &AlibabaseakingauthmachineapiAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaseakingauthmachineapiAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.authmachineapi"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaseakingauthmachineapiAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaseakingauthmachineapiAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdentifyType is IdentifyType Setter
// erp名称
func (r *AlibabaseakingauthmachineapiAPIRequest) SetIdentifyType(_identifyType string) error {
	r._identifyType = _identifyType
	r.Set("identify_type", _identifyType)
	return nil
}

// GetIdentifyType IdentifyType Getter
func (r AlibabaseakingauthmachineapiAPIRequest) GetIdentifyType() string {
	return r._identifyType
}

// SetIdentifier is Identifier Setter
// erp用户id
func (r *AlibabaseakingauthmachineapiAPIRequest) SetIdentifier(_identifier string) error {
	r._identifier = _identifier
	r.Set("identifier", _identifier)
	return nil
}

// GetIdentifier Identifier Getter
func (r AlibabaseakingauthmachineapiAPIRequest) GetIdentifier() string {
	return r._identifier
}

// SetSubIdentifyType is SubIdentifyType Setter
// 店铺所在平台
func (r *AlibabaseakingauthmachineapiAPIRequest) SetSubIdentifyType(_subIdentifyType string) error {
	r._subIdentifyType = _subIdentifyType
	r.Set("sub_identify_type", _subIdentifyType)
	return nil
}

// GetSubIdentifyType SubIdentifyType Getter
func (r AlibabaseakingauthmachineapiAPIRequest) GetSubIdentifyType() string {
	return r._subIdentifyType
}

// SetSubIdentifier is SubIdentifier Setter
// 店铺id(ae为cn开头的店铺id, lazada为邮箱)
func (r *AlibabaseakingauthmachineapiAPIRequest) SetSubIdentifier(_subIdentifier string) error {
	r._subIdentifier = _subIdentifier
	r.Set("sub_identifier", _subIdentifier)
	return nil
}

// GetSubIdentifier SubIdentifier Getter
func (r AlibabaseakingauthmachineapiAPIRequest) GetSubIdentifier() string {
	return r._subIdentifier
}
