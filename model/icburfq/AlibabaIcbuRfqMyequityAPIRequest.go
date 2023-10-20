package icburfq

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicburfqmyequityAPIRequest 我的权益 API请求
// alibaba.icbu.rfq.myequity
//
// 查询供应商权益接口
type AlibabaicburfqmyequityAPIRequest struct {
	model.Params
}

// NewAlibabaicburfqmyequityRequest 初始化AlibabaicburfqmyequityAPIRequest对象
func NewAlibabaicburfqmyequityRequest() *AlibabaicburfqmyequityAPIRequest {
	return &AlibabaicburfqmyequityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicburfqmyequityAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.rfq.myequity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicburfqmyequityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicburfqmyequityAPIRequest) GetRawParams() model.Params {
	return r.Params
}
