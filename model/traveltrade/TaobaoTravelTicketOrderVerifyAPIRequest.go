package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTravelTicketOrderVerifyAPIRequest 飞猪门票核销通知 API请求
// taobao.travel.ticket.order.verify
//
// 系统商通过TOP接口调用通知飞猪门票核销情况
type TaobaoTravelTicketOrderVerifyAPIRequest struct {
	model.Params
	// 使用凭证信息
	_voucherInfos []VoucherInfoDto
	// 外部订单ID
	_outOrderId string
	// (新接入使用voucher_infos)用户短信会收到的确认号
	_confirmCode string
	// 核销次数
	_checkNum int64
	// 下单订单ID
	_orderId int64
	// 门票取消数量
	_returnNum int64
	// 门票总共允许核销次数
	_totalNum int64
	// 供应商核销回调类型：0表示使用本次核销数量（常规），1表示使用总核销数量（已使用+本次）
	_writeOffType int64
}

// NewTaobaoTravelTicketOrderVerifyRequest 初始化TaobaoTravelTicketOrderVerifyAPIRequest对象
func NewTaobaoTravelTicketOrderVerifyRequest() *TaobaoTravelTicketOrderVerifyAPIRequest {
	return &TaobaoTravelTicketOrderVerifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTravelTicketOrderVerifyAPIRequest) GetApiMethodName() string {
	return "taobao.travel.ticket.order.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTravelTicketOrderVerifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTravelTicketOrderVerifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVoucherInfos is VoucherInfos Setter
// 使用凭证信息
func (r *TaobaoTravelTicketOrderVerifyAPIRequest) SetVoucherInfos(_voucherInfos []VoucherInfoDto) error {
	r._voucherInfos = _voucherInfos
	r.Set("voucher_infos", _voucherInfos)
	return nil
}

// GetVoucherInfos VoucherInfos Getter
func (r TaobaoTravelTicketOrderVerifyAPIRequest) GetVoucherInfos() []VoucherInfoDto {
	return r._voucherInfos
}

// SetOutOrderId is OutOrderId Setter
// 外部订单ID
func (r *TaobaoTravelTicketOrderVerifyAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r TaobaoTravelTicketOrderVerifyAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetConfirmCode is ConfirmCode Setter
// (新接入使用voucher_infos)用户短信会收到的确认号
func (r *TaobaoTravelTicketOrderVerifyAPIRequest) SetConfirmCode(_confirmCode string) error {
	r._confirmCode = _confirmCode
	r.Set("confirm_code", _confirmCode)
	return nil
}

// GetConfirmCode ConfirmCode Getter
func (r TaobaoTravelTicketOrderVerifyAPIRequest) GetConfirmCode() string {
	return r._confirmCode
}

// SetCheckNum is CheckNum Setter
// 核销次数
func (r *TaobaoTravelTicketOrderVerifyAPIRequest) SetCheckNum(_checkNum int64) error {
	r._checkNum = _checkNum
	r.Set("check_num", _checkNum)
	return nil
}

// GetCheckNum CheckNum Getter
func (r TaobaoTravelTicketOrderVerifyAPIRequest) GetCheckNum() int64 {
	return r._checkNum
}

// SetOrderId is OrderId Setter
// 下单订单ID
func (r *TaobaoTravelTicketOrderVerifyAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoTravelTicketOrderVerifyAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetReturnNum is ReturnNum Setter
// 门票取消数量
func (r *TaobaoTravelTicketOrderVerifyAPIRequest) SetReturnNum(_returnNum int64) error {
	r._returnNum = _returnNum
	r.Set("return_num", _returnNum)
	return nil
}

// GetReturnNum ReturnNum Getter
func (r TaobaoTravelTicketOrderVerifyAPIRequest) GetReturnNum() int64 {
	return r._returnNum
}

// SetTotalNum is TotalNum Setter
// 门票总共允许核销次数
func (r *TaobaoTravelTicketOrderVerifyAPIRequest) SetTotalNum(_totalNum int64) error {
	r._totalNum = _totalNum
	r.Set("total_num", _totalNum)
	return nil
}

// GetTotalNum TotalNum Getter
func (r TaobaoTravelTicketOrderVerifyAPIRequest) GetTotalNum() int64 {
	return r._totalNum
}

// SetWriteOffType is WriteOffType Setter
// 供应商核销回调类型：0表示使用本次核销数量（常规），1表示使用总核销数量（已使用+本次）
func (r *TaobaoTravelTicketOrderVerifyAPIRequest) SetWriteOffType(_writeOffType int64) error {
	r._writeOffType = _writeOffType
	r.Set("write_off_type", _writeOffType)
	return nil
}

// GetWriteOffType WriteOffType Getter
func (r TaobaoTravelTicketOrderVerifyAPIRequest) GetWriteOffType() int64 {
	return r._writeOffType
}
