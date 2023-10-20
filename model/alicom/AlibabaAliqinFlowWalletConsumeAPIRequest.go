package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowWalletConsumeAPIRequest 流量扣减 API请求
// alibaba.aliqin.flow.wallet.consume
//
// 流量钱包流量扣减接口
type AlibabaAliqinFlowWalletConsumeAPIRequest struct {
	model.Params
	// 扣减流水号
	_serialNo string
	// 扣减原因
	_reason string
	// 备注
	_remark string
	// 扣减流量值
	_flow int64
}

// NewAlibabaAliqinFlowWalletConsumeRequest 初始化AlibabaAliqinFlowWalletConsumeAPIRequest对象
func NewAlibabaAliqinFlowWalletConsumeRequest() *AlibabaAliqinFlowWalletConsumeAPIRequest {
	return &AlibabaAliqinFlowWalletConsumeAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinFlowWalletConsumeAPIRequest) Reset() {
	r._serialNo = ""
	r._reason = ""
	r._remark = ""
	r._flow = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowWalletConsumeAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.wallet.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowWalletConsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFlowWalletConsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSerialNo is SerialNo Setter
// 扣减流水号
func (r *AlibabaAliqinFlowWalletConsumeAPIRequest) SetSerialNo(_serialNo string) error {
	r._serialNo = _serialNo
	r.Set("serial_no", _serialNo)
	return nil
}

// GetSerialNo SerialNo Getter
func (r AlibabaAliqinFlowWalletConsumeAPIRequest) GetSerialNo() string {
	return r._serialNo
}

// SetReason is Reason Setter
// 扣减原因
func (r *AlibabaAliqinFlowWalletConsumeAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r AlibabaAliqinFlowWalletConsumeAPIRequest) GetReason() string {
	return r._reason
}

// SetRemark is Remark Setter
// 备注
func (r *AlibabaAliqinFlowWalletConsumeAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r AlibabaAliqinFlowWalletConsumeAPIRequest) GetRemark() string {
	return r._remark
}

// SetFlow is Flow Setter
// 扣减流量值
func (r *AlibabaAliqinFlowWalletConsumeAPIRequest) SetFlow(_flow int64) error {
	r._flow = _flow
	r.Set("flow", _flow)
	return nil
}

// GetFlow Flow Getter
func (r AlibabaAliqinFlowWalletConsumeAPIRequest) GetFlow() int64 {
	return r._flow
}

var poolAlibabaAliqinFlowWalletConsumeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinFlowWalletConsumeRequest()
	},
}

// GetAlibabaAliqinFlowWalletConsumeRequest 从 sync.Pool 获取 AlibabaAliqinFlowWalletConsumeAPIRequest
func GetAlibabaAliqinFlowWalletConsumeAPIRequest() *AlibabaAliqinFlowWalletConsumeAPIRequest {
	return poolAlibabaAliqinFlowWalletConsumeAPIRequest.Get().(*AlibabaAliqinFlowWalletConsumeAPIRequest)
}

// ReleaseAlibabaAliqinFlowWalletConsumeAPIRequest 将 AlibabaAliqinFlowWalletConsumeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinFlowWalletConsumeAPIRequest(v *AlibabaAliqinFlowWalletConsumeAPIRequest) {
	v.Reset()
	poolAlibabaAliqinFlowWalletConsumeAPIRequest.Put(v)
}
