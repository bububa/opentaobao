package icburfq

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuRfqMyequityAPIRequest 我的权益 API请求
// alibaba.icbu.rfq.myequity
//
// 查询供应商权益接口
type AlibabaIcbuRfqMyequityAPIRequest struct {
	model.Params
}

// NewAlibabaIcbuRfqMyequityRequest 初始化AlibabaIcbuRfqMyequityAPIRequest对象
func NewAlibabaIcbuRfqMyequityRequest() *AlibabaIcbuRfqMyequityAPIRequest {
	return &AlibabaIcbuRfqMyequityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuRfqMyequityAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.rfq.myequity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuRfqMyequityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
