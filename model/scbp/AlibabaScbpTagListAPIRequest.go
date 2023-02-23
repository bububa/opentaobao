package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTagListAPIRequest 查询所有分组 API请求
// alibaba.scbp.tag.list
//
// 查询所有分组
type AlibabaScbpTagListAPIRequest struct {
	model.Params
}

// NewAlibabaScbpTagListRequest 初始化AlibabaScbpTagListAPIRequest对象
func NewAlibabaScbpTagListRequest() *AlibabaScbpTagListAPIRequest {
	return &AlibabaScbpTagListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTagListAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.tag.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTagListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpTagListAPIRequest) GetRawParams() model.Params {
	return r.Params
}
