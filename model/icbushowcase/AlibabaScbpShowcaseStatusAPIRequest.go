package icbushowcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpShowcaseStatusAPIRequest 橱窗状态 API请求
// alibaba.scbp.showcase.status
//
// 查询橱窗状态，如总数、可用数量
type AlibabaScbpShowcaseStatusAPIRequest struct {
	model.Params
}

// NewAlibabaScbpShowcaseStatusRequest 初始化AlibabaScbpShowcaseStatusAPIRequest对象
func NewAlibabaScbpShowcaseStatusRequest() *AlibabaScbpShowcaseStatusAPIRequest {
	return &AlibabaScbpShowcaseStatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpShowcaseStatusAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.showcase.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpShowcaseStatusAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
