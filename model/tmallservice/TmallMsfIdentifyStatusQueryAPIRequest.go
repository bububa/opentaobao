package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMsfIdentifyStatusQueryAPIRequest 喵师傅定案核销状态查询 API请求
// tmall.msf.identify.status.query
//
// 喵师傅定案核销状态查询，供服务商erp系统调用
type TmallMsfIdentifyStatusQueryAPIRequest struct {
	model.Params
	// 天猫订单号
	_orderId int64
	// 服务类型，0 家装的送货上门并安装 1 单向安装 2 建材的送货上门 3 建材的安装
	_serviceType int64
}

// NewTmallMsfIdentifyStatusQueryRequest 初始化TmallMsfIdentifyStatusQueryAPIRequest对象
func NewTmallMsfIdentifyStatusQueryRequest() *TmallMsfIdentifyStatusQueryAPIRequest {
	return &TmallMsfIdentifyStatusQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallMsfIdentifyStatusQueryAPIRequest) Reset() {
	r._orderId = 0
	r._serviceType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMsfIdentifyStatusQueryAPIRequest) GetApiMethodName() string {
	return "tmall.msf.identify.status.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMsfIdentifyStatusQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallMsfIdentifyStatusQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 天猫订单号
func (r *TmallMsfIdentifyStatusQueryAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallMsfIdentifyStatusQueryAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetServiceType is ServiceType Setter
// 服务类型，0 家装的送货上门并安装 1 单向安装 2 建材的送货上门 3 建材的安装
func (r *TmallMsfIdentifyStatusQueryAPIRequest) SetServiceType(_serviceType int64) error {
	r._serviceType = _serviceType
	r.Set("service_type", _serviceType)
	return nil
}

// GetServiceType ServiceType Getter
func (r TmallMsfIdentifyStatusQueryAPIRequest) GetServiceType() int64 {
	return r._serviceType
}

var poolTmallMsfIdentifyStatusQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallMsfIdentifyStatusQueryRequest()
	},
}

// GetTmallMsfIdentifyStatusQueryRequest 从 sync.Pool 获取 TmallMsfIdentifyStatusQueryAPIRequest
func GetTmallMsfIdentifyStatusQueryAPIRequest() *TmallMsfIdentifyStatusQueryAPIRequest {
	return poolTmallMsfIdentifyStatusQueryAPIRequest.Get().(*TmallMsfIdentifyStatusQueryAPIRequest)
}

// ReleaseTmallMsfIdentifyStatusQueryAPIRequest 将 TmallMsfIdentifyStatusQueryAPIRequest 放入 sync.Pool
func ReleaseTmallMsfIdentifyStatusQueryAPIRequest(v *TmallMsfIdentifyStatusQueryAPIRequest) {
	v.Reset()
	poolTmallMsfIdentifyStatusQueryAPIRequest.Put(v)
}
