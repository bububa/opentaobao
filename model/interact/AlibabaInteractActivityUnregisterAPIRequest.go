package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractactivityunregisterAPIRequest ISV互动应用活动关闭服务 API请求
// alibaba.interact.activity.unregister
//
// 卖家在ISV互动应用中设置的活动主动关闭的服务
type AlibabainteractactivityunregisterAPIRequest struct {
	model.Params
	// 互动活动ID
	_bizId string
}

// NewAlibabainteractactivityunregisterRequest 初始化AlibabainteractactivityunregisterAPIRequest对象
func NewAlibabainteractactivityunregisterRequest() *AlibabainteractactivityunregisterAPIRequest {
	return &AlibabainteractactivityunregisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractactivityunregisterAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.activity.unregister"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractactivityunregisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractactivityunregisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizId is BizId Setter
// 互动活动ID
func (r *AlibabainteractactivityunregisterAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// GetBizId BizId Getter
func (r AlibabainteractactivityunregisterAPIRequest) GetBizId() string {
	return r._bizId
}
