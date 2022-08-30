package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleInsurancePaidResultAPIRequest 闲鱼行业保险理赔结果同步 API请求
// alibaba.idle.insurance.paid.result
//
// 闲鱼行业保险理赔结果同步
type AlibabaIdleInsurancePaidResultAPIRequest struct {
	model.Params
	// 理赔id
	_channelOrderId string
	// 理赔款支付id，如支付宝打款id
	_payId string
	// 收款账号id。如支付宝id
	_accountId string
	// 打款失败时错误码
	_payErrCode string
	// 打款失败时错误信息
	_payErrMsg string
	// 打款状态 -1打款失败，1打款成功
	_payStatus int64
	// 打款时间戳，毫秒
	_payTs int64
	// 打款金额，分
	_payCent int64
}

// NewAlibabaIdleInsurancePaidResultRequest 初始化AlibabaIdleInsurancePaidResultAPIRequest对象
func NewAlibabaIdleInsurancePaidResultRequest() *AlibabaIdleInsurancePaidResultAPIRequest {
	return &AlibabaIdleInsurancePaidResultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleInsurancePaidResultAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.insurance.paid.result"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleInsurancePaidResultAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetChannelOrderId is ChannelOrderId Setter
// 理赔id
func (r *AlibabaIdleInsurancePaidResultAPIRequest) SetChannelOrderId(_channelOrderId string) error {
	r._channelOrderId = _channelOrderId
	r.Set("channel_order_id", _channelOrderId)
	return nil
}

// GetChannelOrderId ChannelOrderId Getter
func (r AlibabaIdleInsurancePaidResultAPIRequest) GetChannelOrderId() string {
	return r._channelOrderId
}

// SetPayId is PayId Setter
// 理赔款支付id，如支付宝打款id
func (r *AlibabaIdleInsurancePaidResultAPIRequest) SetPayId(_payId string) error {
	r._payId = _payId
	r.Set("pay_id", _payId)
	return nil
}

// GetPayId PayId Getter
func (r AlibabaIdleInsurancePaidResultAPIRequest) GetPayId() string {
	return r._payId
}

// SetAccountId is AccountId Setter
// 收款账号id。如支付宝id
func (r *AlibabaIdleInsurancePaidResultAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r AlibabaIdleInsurancePaidResultAPIRequest) GetAccountId() string {
	return r._accountId
}

// SetPayErrCode is PayErrCode Setter
// 打款失败时错误码
func (r *AlibabaIdleInsurancePaidResultAPIRequest) SetPayErrCode(_payErrCode string) error {
	r._payErrCode = _payErrCode
	r.Set("pay_err_code", _payErrCode)
	return nil
}

// GetPayErrCode PayErrCode Getter
func (r AlibabaIdleInsurancePaidResultAPIRequest) GetPayErrCode() string {
	return r._payErrCode
}

// SetPayErrMsg is PayErrMsg Setter
// 打款失败时错误信息
func (r *AlibabaIdleInsurancePaidResultAPIRequest) SetPayErrMsg(_payErrMsg string) error {
	r._payErrMsg = _payErrMsg
	r.Set("pay_err_msg", _payErrMsg)
	return nil
}

// GetPayErrMsg PayErrMsg Getter
func (r AlibabaIdleInsurancePaidResultAPIRequest) GetPayErrMsg() string {
	return r._payErrMsg
}

// SetPayStatus is PayStatus Setter
// 打款状态 -1打款失败，1打款成功
func (r *AlibabaIdleInsurancePaidResultAPIRequest) SetPayStatus(_payStatus int64) error {
	r._payStatus = _payStatus
	r.Set("pay_status", _payStatus)
	return nil
}

// GetPayStatus PayStatus Getter
func (r AlibabaIdleInsurancePaidResultAPIRequest) GetPayStatus() int64 {
	return r._payStatus
}

// SetPayTs is PayTs Setter
// 打款时间戳，毫秒
func (r *AlibabaIdleInsurancePaidResultAPIRequest) SetPayTs(_payTs int64) error {
	r._payTs = _payTs
	r.Set("pay_ts", _payTs)
	return nil
}

// GetPayTs PayTs Getter
func (r AlibabaIdleInsurancePaidResultAPIRequest) GetPayTs() int64 {
	return r._payTs
}

// SetPayCent is PayCent Setter
// 打款金额，分
func (r *AlibabaIdleInsurancePaidResultAPIRequest) SetPayCent(_payCent int64) error {
	r._payCent = _payCent
	r.Set("pay_cent", _payCent)
	return nil
}

// GetPayCent PayCent Getter
func (r AlibabaIdleInsurancePaidResultAPIRequest) GetPayCent() int64 {
	return r._payCent
}
