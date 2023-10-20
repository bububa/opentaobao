package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpuopsupplierconsignordershipAPIRequest 履约单商家仓发货结果回传服务 API请求
// alibaba.ascp.uop.supplier.consignorder.ship
//
// ERP通过该接口通知商家仓声明销售订单出库信息,支持履约单纬度全部发货的回传（目前不支持分批回传)
type AlibabaascpuopsupplierconsignordershipAPIRequest struct {
	model.Params
	// 发货回传请求模型
	_consignorderShipRequest *Consignordershiprequest
}

// NewAlibabaascpuopsupplierconsignordershipRequest 初始化AlibabaascpuopsupplierconsignordershipAPIRequest对象
func NewAlibabaascpuopsupplierconsignordershipRequest() *AlibabaascpuopsupplierconsignordershipAPIRequest {
	return &AlibabaascpuopsupplierconsignordershipAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpuopsupplierconsignordershipAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.supplier.consignorder.ship"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpuopsupplierconsignordershipAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpuopsupplierconsignordershipAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConsignorderShipRequest is ConsignorderShipRequest Setter
// 发货回传请求模型
func (r *AlibabaascpuopsupplierconsignordershipAPIRequest) SetConsignorderShipRequest(_consignorderShipRequest *Consignordershiprequest) error {
	r._consignorderShipRequest = _consignorderShipRequest
	r.Set("consignorder_ship_request", _consignorderShipRequest)
	return nil
}

// GetConsignorderShipRequest ConsignorderShipRequest Getter
func (r AlibabaascpuopsupplierconsignordershipAPIRequest) GetConsignorderShipRequest() *Consignordershiprequest {
	return r._consignorderShipRequest
}
