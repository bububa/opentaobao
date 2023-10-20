package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorSocialAPIRequest 社交组件 API请求
// alibaba.interact.sensor.social
//
// 赞，评论 ，关注 新增接口
type AlibabaInteractSensorSocialAPIRequest struct {
	model.Params
	// 系统自动生成
	_id string
}

// NewAlibabaInteractSensorSocialRequest 初始化AlibabaInteractSensorSocialAPIRequest对象
func NewAlibabaInteractSensorSocialRequest() *AlibabaInteractSensorSocialAPIRequest {
	return &AlibabaInteractSensorSocialAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorSocialAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.social"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorSocialAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractSensorSocialAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 系统自动生成
func (r *AlibabaInteractSensorSocialAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaInteractSensorSocialAPIRequest) GetId() string {
	return r._id
}
