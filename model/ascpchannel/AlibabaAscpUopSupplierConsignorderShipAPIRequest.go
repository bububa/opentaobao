package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopSupplierConsignorderShipAPIRequest 履约单商家仓发货结果回传服务 API请求
// alibaba.ascp.uop.supplier.consignorder.ship
//
// ERP通过该接口通知商家仓声明销售订单出库信息,支持履约单纬度全部发货的回传（目前不支持分批回传)
type AlibabaAscpUopSupplierConsignorderShipAPIRequest struct {
	model.Params
	// 发货回传请求模型
	_consignorderShipRequest *Consignordershiprequest
}

// NewAlibabaAscpUopSupplierConsignorderShipRequest 初始化AlibabaAscpUopSupplierConsignorderShipAPIRequest对象
func NewAlibabaAscpUopSupplierConsignorderShipRequest() *AlibabaAscpUopSupplierConsignorderShipAPIRequest {
	return &AlibabaAscpUopSupplierConsignorderShipAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpUopSupplierConsignorderShipAPIRequest) Reset() {
	r._consignorderShipRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierConsignorderShipAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.supplier.consignorder.ship"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierConsignorderShipAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpUopSupplierConsignorderShipAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConsignorderShipRequest is ConsignorderShipRequest Setter
// 发货回传请求模型
func (r *AlibabaAscpUopSupplierConsignorderShipAPIRequest) SetConsignorderShipRequest(_consignorderShipRequest *Consignordershiprequest) error {
	r._consignorderShipRequest = _consignorderShipRequest
	r.Set("consignorder_ship_request", _consignorderShipRequest)
	return nil
}

// GetConsignorderShipRequest ConsignorderShipRequest Getter
func (r AlibabaAscpUopSupplierConsignorderShipAPIRequest) GetConsignorderShipRequest() *Consignordershiprequest {
	return r._consignorderShipRequest
}

var poolAlibabaAscpUopSupplierConsignorderShipAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpUopSupplierConsignorderShipRequest()
	},
}

// GetAlibabaAscpUopSupplierConsignorderShipRequest 从 sync.Pool 获取 AlibabaAscpUopSupplierConsignorderShipAPIRequest
func GetAlibabaAscpUopSupplierConsignorderShipAPIRequest() *AlibabaAscpUopSupplierConsignorderShipAPIRequest {
	return poolAlibabaAscpUopSupplierConsignorderShipAPIRequest.Get().(*AlibabaAscpUopSupplierConsignorderShipAPIRequest)
}

// ReleaseAlibabaAscpUopSupplierConsignorderShipAPIRequest 将 AlibabaAscpUopSupplierConsignorderShipAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpUopSupplierConsignorderShipAPIRequest(v *AlibabaAscpUopSupplierConsignorderShipAPIRequest) {
	v.Reset()
	poolAlibabaAscpUopSupplierConsignorderShipAPIRequest.Put(v)
}
