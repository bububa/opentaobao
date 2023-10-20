package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensormakeupAPIRequest 美妆虚拟试装 API请求
// alibaba.interact.sensor.makeup
//
// 手机淘宝美妆类目虚拟试妆权限，客户端能力（JS－API）
type AlibabainteractsensormakeupAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensormakeupRequest 初始化AlibabainteractsensormakeupAPIRequest对象
func NewAlibabainteractsensormakeupRequest() *AlibabainteractsensormakeupAPIRequest {
	return &AlibabainteractsensormakeupAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensormakeupAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.makeup"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensormakeupAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensormakeupAPIRequest) GetRawParams() model.Params {
	return r.Params
}
