package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorsocialAPIRequest 社交组件 API请求
// alibaba.interact.sensor.social
//
// 赞，评论 ，关注 新增接口
type AlibabainteractsensorsocialAPIRequest struct {
	model.Params
	// 系统自动生成
	_id string
}

// NewAlibabainteractsensorsocialRequest 初始化AlibabainteractsensorsocialAPIRequest对象
func NewAlibabainteractsensorsocialRequest() *AlibabainteractsensorsocialAPIRequest {
	return &AlibabainteractsensorsocialAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensorsocialAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.social"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensorsocialAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensorsocialAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 系统自动生成
func (r *AlibabainteractsensorsocialAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabainteractsensorsocialAPIRequest) GetId() string {
	return r._id
}
