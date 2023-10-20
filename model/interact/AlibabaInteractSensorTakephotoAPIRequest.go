package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensortakephotoAPIRequest takePhoto API请求
// alibaba.interact.sensor.takephoto
//
// 客户端takePhoto
type AlibabainteractsensortakephotoAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensortakephotoRequest 初始化AlibabainteractsensortakephotoAPIRequest对象
func NewAlibabainteractsensortakephotoRequest() *AlibabainteractsensortakephotoAPIRequest {
	return &AlibabainteractsensortakephotoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensortakephotoAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.takephoto"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensortakephotoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensortakephotoAPIRequest) GetRawParams() model.Params {
	return r.Params
}
