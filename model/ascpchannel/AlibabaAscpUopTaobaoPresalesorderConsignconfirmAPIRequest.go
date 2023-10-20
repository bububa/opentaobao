package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest 预售商家仓出库 API请求
// alibaba.ascp.uop.taobao.presalesorder.consignconfirm
//
// 预售商家仓出库
type AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest struct {
	model.Params
	// 预售订单商家仓出库对象
	_presalesOrderConsignConfirmRequest *Presalesorderconsignconfirmrequest
}

// NewAlibabaAscpUopTaobaoPresalesorderConsignconfirmRequest 初始化AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest对象
func NewAlibabaAscpUopTaobaoPresalesorderConsignconfirmRequest() *AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest {
	return &AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest) Reset() {
	r._presalesOrderConsignConfirmRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.taobao.presalesorder.consignconfirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPresalesOrderConsignConfirmRequest is PresalesOrderConsignConfirmRequest Setter
// 预售订单商家仓出库对象
func (r *AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest) SetPresalesOrderConsignConfirmRequest(_presalesOrderConsignConfirmRequest *Presalesorderconsignconfirmrequest) error {
	r._presalesOrderConsignConfirmRequest = _presalesOrderConsignConfirmRequest
	r.Set("presales_order_consign_confirm_request", _presalesOrderConsignConfirmRequest)
	return nil
}

// GetPresalesOrderConsignConfirmRequest PresalesOrderConsignConfirmRequest Getter
func (r AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest) GetPresalesOrderConsignConfirmRequest() *Presalesorderconsignconfirmrequest {
	return r._presalesOrderConsignConfirmRequest
}

var poolAlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpUopTaobaoPresalesorderConsignconfirmRequest()
	},
}

// GetAlibabaAscpUopTaobaoPresalesorderConsignconfirmRequest 从 sync.Pool 获取 AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest
func GetAlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest() *AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest {
	return poolAlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest.Get().(*AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest)
}

// ReleaseAlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest 将 AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest(v *AlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest) {
	v.Reset()
	poolAlibabaAscpUopTaobaoPresalesorderConsignconfirmAPIRequest.Put(v)
}
