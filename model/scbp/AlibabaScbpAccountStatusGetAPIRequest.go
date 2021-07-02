package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAccountStatusGetAPIRequest 查询账户级别关键词推广状态 API请求
// alibaba.scbp.account.status.get
//
// 查询账户级别关键词推广状态
type AlibabaScbpAccountStatusGetAPIRequest struct {
	model.Params
}

// NewAlibabaScbpAccountStatusGetRequest 初始化AlibabaScbpAccountStatusGetAPIRequest对象
func NewAlibabaScbpAccountStatusGetRequest() *AlibabaScbpAccountStatusGetAPIRequest {
	return &AlibabaScbpAccountStatusGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAccountStatusGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.account.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAccountStatusGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
