package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleConsignmentiiOrderGetAPIRequest 闲鱼寄卖V2订单查询 API请求
// alibaba.idle.consignmentii.order.get
//
// 闲鱼寄卖V2服务商以闲鱼交易买家身份查询订单信息
type AlibabaIdleConsignmentiiOrderGetAPIRequest struct {
	model.Params
	// 闲鱼订单ID
	_bizOrderId int64
}

// NewAlibabaIdleConsignmentiiOrderGetRequest 初始化AlibabaIdleConsignmentiiOrderGetAPIRequest对象
func NewAlibabaIdleConsignmentiiOrderGetRequest() *AlibabaIdleConsignmentiiOrderGetAPIRequest {
	return &AlibabaIdleConsignmentiiOrderGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleConsignmentiiOrderGetAPIRequest) Reset() {
	r._bizOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleConsignmentiiOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.consignmentii.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleConsignmentiiOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleConsignmentiiOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 闲鱼订单ID
func (r *AlibabaIdleConsignmentiiOrderGetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r AlibabaIdleConsignmentiiOrderGetAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

var poolAlibabaIdleConsignmentiiOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleConsignmentiiOrderGetRequest()
	},
}

// GetAlibabaIdleConsignmentiiOrderGetRequest 从 sync.Pool 获取 AlibabaIdleConsignmentiiOrderGetAPIRequest
func GetAlibabaIdleConsignmentiiOrderGetAPIRequest() *AlibabaIdleConsignmentiiOrderGetAPIRequest {
	return poolAlibabaIdleConsignmentiiOrderGetAPIRequest.Get().(*AlibabaIdleConsignmentiiOrderGetAPIRequest)
}

// ReleaseAlibabaIdleConsignmentiiOrderGetAPIRequest 将 AlibabaIdleConsignmentiiOrderGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleConsignmentiiOrderGetAPIRequest(v *AlibabaIdleConsignmentiiOrderGetAPIRequest) {
	v.Reset()
	poolAlibabaIdleConsignmentiiOrderGetAPIRequest.Put(v)
}
