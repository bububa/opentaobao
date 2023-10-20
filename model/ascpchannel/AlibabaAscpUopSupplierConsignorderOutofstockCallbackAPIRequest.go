package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest 履约单纬度的仓缺货回告服务 API请求
// alibaba.ascp.uop.supplier.consignorder.outofstock.callback
//
// 商家仓履约单纬度的仓缺货回告接口
type AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest struct {
	model.Params
	// 缺货回告请求模型
	_consignorderOutofstockCallbackRequest *Consignorderoutofstockcallbackrequest
}

// NewAlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest 初始化AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest对象
func NewAlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest() *AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest {
	return &AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest) Reset() {
	r._consignorderOutofstockCallbackRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.supplier.consignorder.outofstock.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConsignorderOutofstockCallbackRequest is ConsignorderOutofstockCallbackRequest Setter
// 缺货回告请求模型
func (r *AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest) SetConsignorderOutofstockCallbackRequest(_consignorderOutofstockCallbackRequest *Consignorderoutofstockcallbackrequest) error {
	r._consignorderOutofstockCallbackRequest = _consignorderOutofstockCallbackRequest
	r.Set("consignorder_outofstock_callback_request", _consignorderOutofstockCallbackRequest)
	return nil
}

// GetConsignorderOutofstockCallbackRequest ConsignorderOutofstockCallbackRequest Getter
func (r AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest) GetConsignorderOutofstockCallbackRequest() *Consignorderoutofstockcallbackrequest {
	return r._consignorderOutofstockCallbackRequest
}

var poolAlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest()
	},
}

// GetAlibabaAscpUopSupplierConsignorderOutofstockCallbackRequest 从 sync.Pool 获取 AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest
func GetAlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest() *AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest {
	return poolAlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest.Get().(*AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest)
}

// ReleaseAlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest 将 AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest(v *AlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest) {
	v.Reset()
	poolAlibabaAscpUopSupplierConsignorderOutofstockCallbackAPIRequest.Put(v)
}
