package ascpchannel

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierConsignorderShipAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.supplier.consignorder.ship"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierConsignorderShipAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
