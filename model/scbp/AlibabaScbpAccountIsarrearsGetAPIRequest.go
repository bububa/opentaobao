package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAccountIsarrearsGetAPIRequest 查询关键词推广账户是否欠款 API请求
// alibaba.scbp.account.isarrears.get
//
// 查询关键词推广账户是否欠款
type AlibabaScbpAccountIsarrearsGetAPIRequest struct {
	model.Params
}

// NewAlibabaScbpAccountIsarrearsGetRequest 初始化AlibabaScbpAccountIsarrearsGetAPIRequest对象
func NewAlibabaScbpAccountIsarrearsGetRequest() *AlibabaScbpAccountIsarrearsGetAPIRequest {
	return &AlibabaScbpAccountIsarrearsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAccountIsarrearsGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.account.isarrears.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAccountIsarrearsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
