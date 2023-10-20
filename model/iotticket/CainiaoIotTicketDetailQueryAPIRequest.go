package iotticket

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoIotTicketDetailQueryAPIRequest IoT售后工单详情查询 API请求
// cainiao.iot.ticket.detail.query
//
// Iot售后工单详情信息查询
type CainiaoIotTicketDetailQueryAPIRequest struct {
	model.Params
	// 服务商唯一编码
	_spCode string
	// 工单Id
	_ticketId int64
}

// NewCainiaoIotTicketDetailQueryRequest 初始化CainiaoIotTicketDetailQueryAPIRequest对象
func NewCainiaoIotTicketDetailQueryRequest() *CainiaoIotTicketDetailQueryAPIRequest {
	return &CainiaoIotTicketDetailQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoIotTicketDetailQueryAPIRequest) Reset() {
	r._spCode = ""
	r._ticketId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketDetailQueryAPIRequest) GetApiMethodName() string {
	return "cainiao.iot.ticket.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketDetailQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoIotTicketDetailQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSpCode is SpCode Setter
// 服务商唯一编码
func (r *CainiaoIotTicketDetailQueryAPIRequest) SetSpCode(_spCode string) error {
	r._spCode = _spCode
	r.Set("sp_code", _spCode)
	return nil
}

// GetSpCode SpCode Getter
func (r CainiaoIotTicketDetailQueryAPIRequest) GetSpCode() string {
	return r._spCode
}

// SetTicketId is TicketId Setter
// 工单Id
func (r *CainiaoIotTicketDetailQueryAPIRequest) SetTicketId(_ticketId int64) error {
	r._ticketId = _ticketId
	r.Set("ticket_id", _ticketId)
	return nil
}

// GetTicketId TicketId Getter
func (r CainiaoIotTicketDetailQueryAPIRequest) GetTicketId() int64 {
	return r._ticketId
}

var poolCainiaoIotTicketDetailQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoIotTicketDetailQueryRequest()
	},
}

// GetCainiaoIotTicketDetailQueryRequest 从 sync.Pool 获取 CainiaoIotTicketDetailQueryAPIRequest
func GetCainiaoIotTicketDetailQueryAPIRequest() *CainiaoIotTicketDetailQueryAPIRequest {
	return poolCainiaoIotTicketDetailQueryAPIRequest.Get().(*CainiaoIotTicketDetailQueryAPIRequest)
}

// ReleaseCainiaoIotTicketDetailQueryAPIRequest 将 CainiaoIotTicketDetailQueryAPIRequest 放入 sync.Pool
func ReleaseCainiaoIotTicketDetailQueryAPIRequest(v *CainiaoIotTicketDetailQueryAPIRequest) {
	v.Reset()
	poolCainiaoIotTicketDetailQueryAPIRequest.Put(v)
}
