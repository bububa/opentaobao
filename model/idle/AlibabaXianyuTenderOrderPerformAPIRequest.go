package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaXianyuTenderOrderPerformAPIRequest 闲鱼暗拍订单履约 API请求
// alibaba.xianyu.tender.order.perform
//
// 闲鱼暗拍订单履约
type AlibabaXianyuTenderOrderPerformAPIRequest struct {
	model.Params
	// 入参
	_param0 *TenderOrderSynDto
}

// NewAlibabaXianyuTenderOrderPerformRequest 初始化AlibabaXianyuTenderOrderPerformAPIRequest对象
func NewAlibabaXianyuTenderOrderPerformRequest() *AlibabaXianyuTenderOrderPerformAPIRequest {
	return &AlibabaXianyuTenderOrderPerformAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaXianyuTenderOrderPerformAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaXianyuTenderOrderPerformAPIRequest) GetApiMethodName() string {
	return "alibaba.xianyu.tender.order.perform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaXianyuTenderOrderPerformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaXianyuTenderOrderPerformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *AlibabaXianyuTenderOrderPerformAPIRequest) SetParam0(_param0 *TenderOrderSynDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaXianyuTenderOrderPerformAPIRequest) GetParam0() *TenderOrderSynDto {
	return r._param0
}

var poolAlibabaXianyuTenderOrderPerformAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaXianyuTenderOrderPerformRequest()
	},
}

// GetAlibabaXianyuTenderOrderPerformRequest 从 sync.Pool 获取 AlibabaXianyuTenderOrderPerformAPIRequest
func GetAlibabaXianyuTenderOrderPerformAPIRequest() *AlibabaXianyuTenderOrderPerformAPIRequest {
	return poolAlibabaXianyuTenderOrderPerformAPIRequest.Get().(*AlibabaXianyuTenderOrderPerformAPIRequest)
}

// ReleaseAlibabaXianyuTenderOrderPerformAPIRequest 将 AlibabaXianyuTenderOrderPerformAPIRequest 放入 sync.Pool
func ReleaseAlibabaXianyuTenderOrderPerformAPIRequest(v *AlibabaXianyuTenderOrderPerformAPIRequest) {
	v.Reset()
	poolAlibabaXianyuTenderOrderPerformAPIRequest.Put(v)
}
