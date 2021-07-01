package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFlowWalletConsumeAPIRequest
流量扣减 API请求
alibaba.aliqin.flow.wallet.consume

流量钱包流量扣减接口 */
type AlibabaAliqinFlowWalletConsumeAPIRequest struct {
	model.Params
	// 扣减流量值
	_flow int64
	// 扣减流水号
	_serialNo string
	// 扣减原因
	_reason string
	// 备注
	_remark string
}

// NewAlibabaAliqinFlowWalletConsumeRequest 初始化AlibabaAliqinFlowWalletConsumeAPIRequest对象
func NewAlibabaAliqinFlowWalletConsumeRequest() *AlibabaAliqinFlowWalletConsumeAPIRequest {
	return &AlibabaAliqinFlowWalletConsumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowWalletConsumeAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.wallet.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowWalletConsumeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Flow Setter
// 扣减流量值
func (r *AlibabaAliqinFlowWalletConsumeAPIRequest) SetFlow(_flow int64) error {
	r._flow = _flow
	r.Set("flow", _flow)
	return nil
}

// Get Flow Getter
func (r AlibabaAliqinFlowWalletConsumeAPIRequest) GetFlow() int64 {
	return r._flow
}

// Set is SerialNo Setter
// 扣减流水号
func (r *AlibabaAliqinFlowWalletConsumeAPIRequest) SetSerialNo(_serialNo string) error {
	r._serialNo = _serialNo
	r.Set("serial_no", _serialNo)
	return nil
}

// Get SerialNo Getter
func (r AlibabaAliqinFlowWalletConsumeAPIRequest) GetSerialNo() string {
	return r._serialNo
}

// Set is Reason Setter
// 扣减原因
func (r *AlibabaAliqinFlowWalletConsumeAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// Get Reason Getter
func (r AlibabaAliqinFlowWalletConsumeAPIRequest) GetReason() string {
	return r._reason
}

// Set is Remark Setter
// 备注
func (r *AlibabaAliqinFlowWalletConsumeAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// Get Remark Getter
func (r AlibabaAliqinFlowWalletConsumeAPIRequest) GetRemark() string {
	return r._remark
}
