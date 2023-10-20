package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitCaseGetAPIRequest 获取案件信息接口v2版本 API请求
// alibaba.legal.suit.case.get
//
// 获取案件信息
type AlibabaLegalSuitCaseGetAPIRequest struct {
	model.Params
	// 案件id
	_id int64
}

// NewAlibabaLegalSuitCaseGetRequest 初始化AlibabaLegalSuitCaseGetAPIRequest对象
func NewAlibabaLegalSuitCaseGetRequest() *AlibabaLegalSuitCaseGetAPIRequest {
	return &AlibabaLegalSuitCaseGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitCaseGetAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.case.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitCaseGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalSuitCaseGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 案件id
func (r *AlibabaLegalSuitCaseGetAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaLegalSuitCaseGetAPIRequest) GetId() int64 {
	return r._id
}
