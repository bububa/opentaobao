package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeAgreementPreshowAPIRequest) Reset() {
	r._id = 0
	r._type = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeAgreementPreshowAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.agreement.preshow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeAgreementPreshowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeAgreementPreshowAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihouseNewhomeAgreementPreshowAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeAgreementPreshowRequest()
	},
}

// GetAlibabaAlihouseNewhomeAgreementPreshowRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeAgreementPreshowAPIRequest
func GetAlibabaAlihouseNewhomeAgreementPreshowAPIRequest() *AlibabaAlihouseNewhomeAgreementPreshowAPIRequest {
	return poolAlibabaAlihouseNewhomeAgreementPreshowAPIRequest.Get().(*AlibabaAlihouseNewhomeAgreementPreshowAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeAgreementPreshowAPIRequest 将 AlibabaAlihouseNewhomeAgreementPreshowAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeAgreementPreshowAPIRequest(v *AlibabaAlihouseNewhomeAgreementPreshowAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeAgreementPreshowAPIRequest.Put(v)
}
