package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleConsignmentOrderGetAPIRequest 闲鱼帮卖订单查询 API请求
// alibaba.idle.consignment.order.get
//
// 闲鱼帮卖服务商以闲鱼交易买家身份查询订单信息
type AlibabaIdleConsignmentOrderGetAPIRequest struct {
	model.Params
	// 闲鱼订单ID
	_bizOrderId int64
}

// NewAlibabaIdleConsignmentOrderGetRequest 初始化AlibabaIdleConsignmentOrderGetAPIRequest对象
func NewAlibabaIdleConsignmentOrderGetRequest() *AlibabaIdleConsignmentOrderGetAPIRequest {
	return &AlibabaIdleConsignmentOrderGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleConsignmentOrderGetAPIRequest) Reset() {
	r._bizOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleConsignmentOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.consignment.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleConsignmentOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleConsignmentOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 闲鱼订单ID
func (r *AlibabaIdleConsignmentOrderGetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r AlibabaIdleConsignmentOrderGetAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

var poolAlibabaIdleConsignmentOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleConsignmentOrderGetRequest()
	},
}

// GetAlibabaIdleConsignmentOrderGetRequest 从 sync.Pool 获取 AlibabaIdleConsignmentOrderGetAPIRequest
func GetAlibabaIdleConsignmentOrderGetAPIRequest() *AlibabaIdleConsignmentOrderGetAPIRequest {
	return poolAlibabaIdleConsignmentOrderGetAPIRequest.Get().(*AlibabaIdleConsignmentOrderGetAPIRequest)
}

// ReleaseAlibabaIdleConsignmentOrderGetAPIRequest 将 AlibabaIdleConsignmentOrderGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleConsignmentOrderGetAPIRequest(v *AlibabaIdleConsignmentOrderGetAPIRequest) {
	v.Reset()
	poolAlibabaIdleConsignmentOrderGetAPIRequest.Put(v)
}
