package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractcurrentmixusernickAPIRequest 手淘混淆nick开放接口鉴权专用 API请求
// alibaba.interact.current.mixusernick
//
// 手淘混淆nick开放接口鉴权专用，无数据输入输出。
type AlibabainteractcurrentmixusernickAPIRequest struct {
	model.Params
}

// NewAlibabainteractcurrentmixusernickRequest 初始化AlibabainteractcurrentmixusernickAPIRequest对象
func NewAlibabainteractcurrentmixusernickRequest() *AlibabainteractcurrentmixusernickAPIRequest {
	return &AlibabainteractcurrentmixusernickAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractcurrentmixusernickAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.current.mixusernick"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractcurrentmixusernickAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractcurrentmixusernickAPIRequest) GetRawParams() model.Params {
	return r.Params
}
