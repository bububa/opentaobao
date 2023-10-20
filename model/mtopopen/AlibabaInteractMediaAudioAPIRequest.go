package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractmediaaudioAPIRequest 音频相关鉴权接口 API请求
// alibaba.interact.media.audio
//
// 新音频包的鉴权接口
type AlibabainteractmediaaudioAPIRequest struct {
	model.Params
	// 系统自动生成
	_id string
}

// NewAlibabainteractmediaaudioRequest 初始化AlibabainteractmediaaudioAPIRequest对象
func NewAlibabainteractmediaaudioRequest() *AlibabainteractmediaaudioAPIRequest {
	return &AlibabainteractmediaaudioAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractmediaaudioAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.media.audio"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractmediaaudioAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractmediaaudioAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 系统自动生成
func (r *AlibabainteractmediaaudioAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabainteractmediaaudioAPIRequest) GetId() string {
	return r._id
}
