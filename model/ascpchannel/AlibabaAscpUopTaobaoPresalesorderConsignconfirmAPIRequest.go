package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest
预售商家仓出库 API请求
alibaba.ascp.uop.taobao.presalesorder.consignconfirm

预售商家仓出库 */
type AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest struct {
	model.Params
	// 预售订单商家仓出库对象
	_presalesOrderConsignConfirmRequest *Presalesorderconsignconfirmrequest
}

// NewAlibabaAscpUopTaobaoPresalesorderConsignconfirmRequest 初始化AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest对象
func NewAlibabaAscpUopTaobaoPresalesorderConsignconfirmRequest() *AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest {
	return &AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.taobao.presalesorder.consignconfirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PresalesOrderConsignConfirmRequest Setter
// 预售订单商家仓出库对象
func (r *AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest) SetPresalesOrderConsignConfirmRequest(_presalesOrderConsignConfirmRequest *Presalesorderconsignconfirmrequest) error {
	r._presalesOrderConsignConfirmRequest = _presalesOrderConsignConfirmRequest
	r.Set("presales_order_consign_confirm_request", _presalesOrderConsignConfirmRequest)
	return nil
}

// Get PresalesOrderConsignConfirmRequest Getter
func (r AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest) GetPresalesOrderConsignConfirmRequest() *Presalesorderconsignconfirmrequest {
	return r._presalesOrderConsignConfirmRequest
}
