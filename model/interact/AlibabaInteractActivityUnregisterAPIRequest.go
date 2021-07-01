package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractActivityUnregisterAPIRequest
ISV互动应用活动关闭服务 API请求
alibaba.interact.activity.unregister

卖家在ISV互动应用中设置的活动主动关闭的服务 */
type AlibabaInteractActivityUnregisterAPIRequest struct {
	model.Params
	// 互动活动ID
	_bizId string
}

// NewAlibabaInteractActivityUnregisterRequest 初始化AlibabaInteractActivityUnregisterAPIRequest对象
func NewAlibabaInteractActivityUnregisterRequest() *AlibabaInteractActivityUnregisterAPIRequest {
	return &AlibabaInteractActivityUnregisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractActivityUnregisterAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.activity.unregister"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractActivityUnregisterAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BizId Setter
// 互动活动ID
func (r *AlibabaInteractActivityUnregisterAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// Get BizId Getter
func (r AlibabaInteractActivityUnregisterAPIRequest) GetBizId() string {
	return r._bizId
}
