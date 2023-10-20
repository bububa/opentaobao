package tblogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsSellerWriteoffAPIRequest 商家配送核销 API请求
// alibaba.ascp.logistics.seller.writeoff
//
// 商家配送核销
type AlibabaAscpLogisticsSellerWriteoffAPIRequest struct {
	model.Params
	// 核销码
	_receiveCode string
	// 所要核销的核销物流单ID，在alibaba.ascp.logistics.seller.orders.get中获取。
	_lpOrderId int64
}

// NewAlibabaAscpLogisticsSellerWriteoffRequest 初始化AlibabaAscpLogisticsSellerWriteoffAPIRequest对象
func NewAlibabaAscpLogisticsSellerWriteoffRequest() *AlibabaAscpLogisticsSellerWriteoffAPIRequest {
	return &AlibabaAscpLogisticsSellerWriteoffAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpLogisticsSellerWriteoffAPIRequest) Reset() {
	r._receiveCode = ""
	r._lpOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpLogisticsSellerWriteoffAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.seller.writeoff"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpLogisticsSellerWriteoffAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpLogisticsSellerWriteoffAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReceiveCode is ReceiveCode Setter
// 核销码
func (r *AlibabaAscpLogisticsSellerWriteoffAPIRequest) SetReceiveCode(_receiveCode string) error {
	r._receiveCode = _receiveCode
	r.Set("receive_code", _receiveCode)
	return nil
}

// GetReceiveCode ReceiveCode Getter
func (r AlibabaAscpLogisticsSellerWriteoffAPIRequest) GetReceiveCode() string {
	return r._receiveCode
}

// SetLpOrderId is LpOrderId Setter
// 所要核销的核销物流单ID，在alibaba.ascp.logistics.seller.orders.get中获取。
func (r *AlibabaAscpLogisticsSellerWriteoffAPIRequest) SetLpOrderId(_lpOrderId int64) error {
	r._lpOrderId = _lpOrderId
	r.Set("lp_order_id", _lpOrderId)
	return nil
}

// GetLpOrderId LpOrderId Getter
func (r AlibabaAscpLogisticsSellerWriteoffAPIRequest) GetLpOrderId() int64 {
	return r._lpOrderId
}

var poolAlibabaAscpLogisticsSellerWriteoffAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpLogisticsSellerWriteoffRequest()
	},
}

// GetAlibabaAscpLogisticsSellerWriteoffRequest 从 sync.Pool 获取 AlibabaAscpLogisticsSellerWriteoffAPIRequest
func GetAlibabaAscpLogisticsSellerWriteoffAPIRequest() *AlibabaAscpLogisticsSellerWriteoffAPIRequest {
	return poolAlibabaAscpLogisticsSellerWriteoffAPIRequest.Get().(*AlibabaAscpLogisticsSellerWriteoffAPIRequest)
}

// ReleaseAlibabaAscpLogisticsSellerWriteoffAPIRequest 将 AlibabaAscpLogisticsSellerWriteoffAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpLogisticsSellerWriteoffAPIRequest(v *AlibabaAscpLogisticsSellerWriteoffAPIRequest) {
	v.Reset()
	poolAlibabaAscpLogisticsSellerWriteoffAPIRequest.Put(v)
}
