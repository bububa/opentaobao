package iotticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoiotticketdetailqueryAPIRequest IoT售后工单详情查询 API请求
// cainiao.iot.ticket.detail.query
//
// Iot售后工单详情信息查询
type CainiaoiotticketdetailqueryAPIRequest struct {
	model.Params
	// 服务商唯一编码
	_spCode string
	// 工单Id
	_ticketId int64
}

// NewCainiaoiotticketdetailqueryRequest 初始化CainiaoiotticketdetailqueryAPIRequest对象
func NewCainiaoiotticketdetailqueryRequest() *CainiaoiotticketdetailqueryAPIRequest {
	return &CainiaoiotticketdetailqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoiotticketdetailqueryAPIRequest) GetApiMethodName() string {
	return "cainiao.iot.ticket.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoiotticketdetailqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoiotticketdetailqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSpCode is SpCode Setter
// 服务商唯一编码
func (r *CainiaoiotticketdetailqueryAPIRequest) SetSpCode(_spCode string) error {
	r._spCode = _spCode
	r.Set("sp_code", _spCode)
	return nil
}

// GetSpCode SpCode Getter
func (r CainiaoiotticketdetailqueryAPIRequest) GetSpCode() string {
	return r._spCode
}

// SetTicketId is TicketId Setter
// 工单Id
func (r *CainiaoiotticketdetailqueryAPIRequest) SetTicketId(_ticketId int64) error {
	r._ticketId = _ticketId
	r.Set("ticket_id", _ticketId)
	return nil
}

// GetTicketId TicketId Getter
func (r CainiaoiotticketdetailqueryAPIRequest) GetTicketId() int64 {
	return r._ticketId
}
