package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractWirelessDrawAPIRequest 双11到店互动无线端抽奖接口鉴权 API请求
// alibaba.interact.wireless.draw
//
// 双11到店互动无线端mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 坯子
type AlibabaInteractWirelessDrawAPIRequest struct {
	model.Params
}

// NewAlibabaInteractWirelessDrawRequest 初始化AlibabaInteractWirelessDrawAPIRequest对象
func NewAlibabaInteractWirelessDrawRequest() *AlibabaInteractWirelessDrawAPIRequest {
	return &AlibabaInteractWirelessDrawAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractWirelessDrawAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractWirelessDrawAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.wireless.draw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractWirelessDrawAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractWirelessDrawAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaInteractWirelessDrawAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractWirelessDrawRequest()
	},
}

// GetAlibabaInteractWirelessDrawRequest 从 sync.Pool 获取 AlibabaInteractWirelessDrawAPIRequest
func GetAlibabaInteractWirelessDrawAPIRequest() *AlibabaInteractWirelessDrawAPIRequest {
	return poolAlibabaInteractWirelessDrawAPIRequest.Get().(*AlibabaInteractWirelessDrawAPIRequest)
}

// ReleaseAlibabaInteractWirelessDrawAPIRequest 将 AlibabaInteractWirelessDrawAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractWirelessDrawAPIRequest(v *AlibabaInteractWirelessDrawAPIRequest) {
	v.Reset()
	poolAlibabaInteractWirelessDrawAPIRequest.Put(v)
}
