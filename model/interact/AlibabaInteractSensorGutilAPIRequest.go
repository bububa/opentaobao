package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorgutilAPIRequest canvas工具包 API请求
// alibaba.interact.sensor.gutil
//
// canvas工具包
type AlibabainteractsensorgutilAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensorgutilRequest 初始化AlibabainteractsensorgutilAPIRequest对象
func NewAlibabainteractsensorgutilRequest() *AlibabainteractsensorgutilAPIRequest {
	return &AlibabainteractsensorgutilAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensorgutilAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.gutil"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensorgutilAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensorgutilAPIRequest) GetRawParams() model.Params {
	return r.Params
}
