package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalcaseentrustgetAPIRequest 委托 API请求
// alibaba.legal.case.entrust.get
//
// 获取委托案件的基本信息
type AlibabalegalcaseentrustgetAPIRequest struct {
	model.Params
	// 委托id
	_entrustId int64
}

// NewAlibabalegalcaseentrustgetRequest 初始化AlibabalegalcaseentrustgetAPIRequest对象
func NewAlibabalegalcaseentrustgetRequest() *AlibabalegalcaseentrustgetAPIRequest {
	return &AlibabalegalcaseentrustgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalcaseentrustgetAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.entrust.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalcaseentrustgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalcaseentrustgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntrustId is EntrustId Setter
// 委托id
func (r *AlibabalegalcaseentrustgetAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabalegalcaseentrustgetAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}
