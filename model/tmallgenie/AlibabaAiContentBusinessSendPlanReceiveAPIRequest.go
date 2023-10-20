package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAiContentBusinessSendPlanReceiveAPIRequest 天猫精灵商业化采销发放计划领取 API请求
// alibaba.ai.content.business.send.plan.receive
//
// 天猫精灵商业化采销发放计划领取
type AlibabaAiContentBusinessSendPlanReceiveAPIRequest struct {
	model.Params
	// 入参对象
	_param0 *PurchaseForOuterDto
}

// NewAlibabaAiContentBusinessSendPlanReceiveRequest 初始化AlibabaAiContentBusinessSendPlanReceiveAPIRequest对象
func NewAlibabaAiContentBusinessSendPlanReceiveRequest() *AlibabaAiContentBusinessSendPlanReceiveAPIRequest {
	return &AlibabaAiContentBusinessSendPlanReceiveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAiContentBusinessSendPlanReceiveAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAiContentBusinessSendPlanReceiveAPIRequest) GetApiMethodName() string {
	return "alibaba.ai.content.business.send.plan.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAiContentBusinessSendPlanReceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAiContentBusinessSendPlanReceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参对象
func (r *AlibabaAiContentBusinessSendPlanReceiveAPIRequest) SetParam0(_param0 *PurchaseForOuterDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaAiContentBusinessSendPlanReceiveAPIRequest) GetParam0() *PurchaseForOuterDto {
	return r._param0
}

var poolAlibabaAiContentBusinessSendPlanReceiveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAiContentBusinessSendPlanReceiveRequest()
	},
}

// GetAlibabaAiContentBusinessSendPlanReceiveRequest 从 sync.Pool 获取 AlibabaAiContentBusinessSendPlanReceiveAPIRequest
func GetAlibabaAiContentBusinessSendPlanReceiveAPIRequest() *AlibabaAiContentBusinessSendPlanReceiveAPIRequest {
	return poolAlibabaAiContentBusinessSendPlanReceiveAPIRequest.Get().(*AlibabaAiContentBusinessSendPlanReceiveAPIRequest)
}

// ReleaseAlibabaAiContentBusinessSendPlanReceiveAPIRequest 将 AlibabaAiContentBusinessSendPlanReceiveAPIRequest 放入 sync.Pool
func ReleaseAlibabaAiContentBusinessSendPlanReceiveAPIRequest(v *AlibabaAiContentBusinessSendPlanReceiveAPIRequest) {
	v.Reset()
	poolAlibabaAiContentBusinessSendPlanReceiveAPIRequest.Put(v)
}
