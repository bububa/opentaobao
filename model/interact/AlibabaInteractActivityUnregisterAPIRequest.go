package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractActivityUnregisterAPIRequest ISV互动应用活动关闭服务 API请求
// alibaba.interact.activity.unregister
//
// 卖家在ISV互动应用中设置的活动主动关闭的服务
type AlibabaInteractActivityUnregisterAPIRequest struct {
	model.Params
	// 互动活动ID
	_bizId string
}

// NewAlibabaInteractActivityUnregisterRequest 初始化AlibabaInteractActivityUnregisterAPIRequest对象
func NewAlibabaInteractActivityUnregisterRequest() *AlibabaInteractActivityUnregisterAPIRequest {
	return &AlibabaInteractActivityUnregisterAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractActivityUnregisterAPIRequest) Reset() {
	r._bizId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractActivityUnregisterAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.activity.unregister"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractActivityUnregisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractActivityUnregisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizId is BizId Setter
// 互动活动ID
func (r *AlibabaInteractActivityUnregisterAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// GetBizId BizId Getter
func (r AlibabaInteractActivityUnregisterAPIRequest) GetBizId() string {
	return r._bizId
}

var poolAlibabaInteractActivityUnregisterAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractActivityUnregisterRequest()
	},
}

// GetAlibabaInteractActivityUnregisterRequest 从 sync.Pool 获取 AlibabaInteractActivityUnregisterAPIRequest
func GetAlibabaInteractActivityUnregisterAPIRequest() *AlibabaInteractActivityUnregisterAPIRequest {
	return poolAlibabaInteractActivityUnregisterAPIRequest.Get().(*AlibabaInteractActivityUnregisterAPIRequest)
}

// ReleaseAlibabaInteractActivityUnregisterAPIRequest 将 AlibabaInteractActivityUnregisterAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractActivityUnregisterAPIRequest(v *AlibabaInteractActivityUnregisterAPIRequest) {
	v.Reset()
	poolAlibabaInteractActivityUnregisterAPIRequest.Put(v)
}
