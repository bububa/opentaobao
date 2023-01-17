package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkMessageConfigListAPIRequest 查询消息开关配置列表 API请求
// alibaba.alink.message.config.list
//
// 阿里智能获取消息开关配置列表
type AlibabaAlinkMessageConfigListAPIRequest struct {
	model.Params
}

// NewAlibabaAlinkMessageConfigListRequest 初始化AlibabaAlinkMessageConfigListAPIRequest对象
func NewAlibabaAlinkMessageConfigListRequest() *AlibabaAlinkMessageConfigListAPIRequest {
	return &AlibabaAlinkMessageConfigListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlinkMessageConfigListAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.message.config.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlinkMessageConfigListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlinkMessageConfigListAPIRequest) GetRawParams() model.Params {
	return r.Params
}
