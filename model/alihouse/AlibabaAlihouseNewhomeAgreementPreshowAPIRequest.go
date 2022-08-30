package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeAgreementPreshowAPIRequest 预览地址获取接口 API请求
// alibaba.alihouse.newhome.agreement.preshow
//
// 预览地址获取接口
type AlibabaAlihouseNewhomeAgreementPreshowAPIRequest struct {
	model.Params
	// 协议ID
	_id int64
	// 1-协议 2-签章
	_type int64
}

// NewAlibabaAlihouseNewhomeAgreementPreshowRequest 初始化AlibabaAlihouseNewhomeAgreementPreshowAPIRequest对象
func NewAlibabaAlihouseNewhomeAgreementPreshowRequest() *AlibabaAlihouseNewhomeAgreementPreshowAPIRequest {
	return &AlibabaAlihouseNewhomeAgreementPreshowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeAgreementPreshowAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.agreement.preshow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeAgreementPreshowAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetId is Id Setter
// 协议ID
func (r *AlibabaAlihouseNewhomeAgreementPreshowAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaAlihouseNewhomeAgreementPreshowAPIRequest) GetId() int64 {
	return r._id
}

// SetType is Type Setter
// 1-协议 2-签章
func (r *AlibabaAlihouseNewhomeAgreementPreshowAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaAlihouseNewhomeAgreementPreshowAPIRequest) GetType() int64 {
	return r._type
}
