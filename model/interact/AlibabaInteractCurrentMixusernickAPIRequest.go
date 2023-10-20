package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractCurrentMixusernickAPIRequest 手淘混淆nick开放接口鉴权专用 API请求
// alibaba.interact.current.mixusernick
//
// 手淘混淆nick开放接口鉴权专用，无数据输入输出。
type AlibabaInteractCurrentMixusernickAPIRequest struct {
	model.Params
}

// NewAlibabaInteractCurrentMixusernickRequest 初始化AlibabaInteractCurrentMixusernickAPIRequest对象
func NewAlibabaInteractCurrentMixusernickRequest() *AlibabaInteractCurrentMixusernickAPIRequest {
	return &AlibabaInteractCurrentMixusernickAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractCurrentMixusernickAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractCurrentMixusernickAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.current.mixusernick"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractCurrentMixusernickAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractCurrentMixusernickAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractCurrentMixusernickAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractCurrentMixusernickRequest()
	},
}

// GetAlibabaInteractCurrentMixusernickRequest 从 sync.Pool 获取 AlibabaInteractCurrentMixusernickAPIRequest
func GetAlibabaInteractCurrentMixusernickAPIRequest() *AlibabaInteractCurrentMixusernickAPIRequest {
	return poolAlibabaInteractCurrentMixusernickAPIRequest.Get().(*AlibabaInteractCurrentMixusernickAPIRequest)
}

// ReleaseAlibabaInteractCurrentMixusernickAPIRequest 将 AlibabaInteractCurrentMixusernickAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractCurrentMixusernickAPIRequest(v *AlibabaInteractCurrentMixusernickAPIRequest) {
	v.Reset()
	poolAlibabaInteractCurrentMixusernickAPIRequest.Put(v)
}
