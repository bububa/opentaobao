package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorglueAPIRequest 视频播放 API请求
// alibaba.interact.sensor.glue
//
// 视频播放
type AlibabainteractsensorglueAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensorglueRequest 初始化AlibabainteractsensorglueAPIRequest对象
func NewAlibabainteractsensorglueRequest() *AlibabainteractsensorglueAPIRequest {
	return &AlibabainteractsensorglueAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensorglueAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.glue"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensorglueAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensorglueAPIRequest) GetRawParams() model.Params {
	return r.Params
}
