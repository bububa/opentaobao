package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest 商家仓物流发货推单接单回告 API请求
// alibaba.ascp.uop.supplier.consignorder.notify.received
//
// ASCP通过该接口接收商家仓开始接单生产订单对应的物流订单信息
type AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest struct {
	model.Params
	// qimen.alibaba.ascp.uop.consignorder.notify报文中的supplierId字段值
	_supplierId string
	// qimen.alibaba.ascp.uop.consignorder.notify报文中bizOrderCode履约单号
	_bizOrderCode string
	// 业务请求时间
	_bizTime string
	// 一盘货业务模式，默认为0代表商家仓商家配，为1代表商家仓自营配 (为1时会强制校验配CP和单号必须与取号时一致，且多包裹必须一次性发货)
	_businessModel string
}

// NewAlibabaAscpUopSupplierConsignorderNotifyReceivedRequest 初始化AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest对象
func NewAlibabaAscpUopSupplierConsignorderNotifyReceivedRequest() *AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest {
	return &AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest) Reset() {
	r._supplierId = ""
	r._bizOrderCode = ""
	r._bizTime = ""
	r._businessModel = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.supplier.consignorder.notify.received"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierId is SupplierId Setter
// qimen.alibaba.ascp.uop.consignorder.notify报文中的supplierId字段值
func (r *AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest) SetSupplierId(_supplierId string) error {
	r._supplierId = _supplierId
	r.Set("supplier_id", _supplierId)
	return nil
}

// GetSupplierId SupplierId Getter
func (r AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest) GetSupplierId() string {
	return r._supplierId
}

// SetBizOrderCode is BizOrderCode Setter
// qimen.alibaba.ascp.uop.consignorder.notify报文中bizOrderCode履约单号
func (r *AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest) SetBizOrderCode(_bizOrderCode string) error {
	r._bizOrderCode = _bizOrderCode
	r.Set("biz_order_code", _bizOrderCode)
	return nil
}

// GetBizOrderCode BizOrderCode Getter
func (r AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest) GetBizOrderCode() string {
	return r._bizOrderCode
}

// SetBizTime is BizTime Setter
// 业务请求时间
func (r *AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest) SetBizTime(_bizTime string) error {
	r._bizTime = _bizTime
	r.Set("biz_time", _bizTime)
	return nil
}

// GetBizTime BizTime Getter
func (r AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest) GetBizTime() string {
	return r._bizTime
}

// SetBusinessModel is BusinessModel Setter
// 一盘货业务模式，默认为0代表商家仓商家配，为1代表商家仓自营配 (为1时会强制校验配CP和单号必须与取号时一致，且多包裹必须一次性发货)
func (r *AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest) SetBusinessModel(_businessModel string) error {
	r._businessModel = _businessModel
	r.Set("business_model", _businessModel)
	return nil
}

// GetBusinessModel BusinessModel Getter
func (r AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest) GetBusinessModel() string {
	return r._businessModel
}

var poolAlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpUopSupplierConsignorderNotifyReceivedRequest()
	},
}

// GetAlibabaAscpUopSupplierConsignorderNotifyReceivedRequest 从 sync.Pool 获取 AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest
func GetAlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest() *AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest {
	return poolAlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest.Get().(*AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest)
}

// ReleaseAlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest 将 AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest(v *AlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest) {
	v.Reset()
	poolAlibabaAscpUopSupplierConsignorderNotifyReceivedAPIRequest.Put(v)
}
