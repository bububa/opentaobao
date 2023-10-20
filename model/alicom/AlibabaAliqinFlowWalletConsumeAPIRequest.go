package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinflowwalletconsumeAPIRequest 流量扣减 API请求
// alibaba.aliqin.flow.wallet.consume
//
// 流量钱包流量扣减接口
type AlibabaaliqinflowwalletconsumeAPIRequest struct {
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

// NewAlibabaaliqinflowwalletconsumeRequest 初始化AlibabaaliqinflowwalletconsumeAPIRequest对象
func NewAlibabaaliqinflowwalletconsumeRequest() *AlibabaaliqinflowwalletconsumeAPIRequest {
	return &AlibabaaliqinflowwalletconsumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinflowwalletconsumeAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.wallet.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinflowwalletconsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinflowwalletconsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSerialNo is SerialNo Setter
// 扣减流水号
func (r *AlibabaaliqinflowwalletconsumeAPIRequest) SetSerialNo(_serialNo string) error {
	r._serialNo = _serialNo
	r.Set("serial_no", _serialNo)
	return nil
}

// GetSerialNo SerialNo Getter
func (r AlibabaaliqinflowwalletconsumeAPIRequest) GetSerialNo() string {
	return r._serialNo
}

// SetReason is Reason Setter
// 扣减原因
func (r *AlibabaaliqinflowwalletconsumeAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r AlibabaaliqinflowwalletconsumeAPIRequest) GetReason() string {
	return r._reason
}

// SetRemark is Remark Setter
// 备注
func (r *AlibabaaliqinflowwalletconsumeAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r AlibabaaliqinflowwalletconsumeAPIRequest) GetRemark() string {
	return r._remark
}

// SetFlow is Flow Setter
// 扣减流量值
func (r *AlibabaaliqinflowwalletconsumeAPIRequest) SetFlow(_flow int64) error {
	r._flow = _flow
	r.Set("flow", _flow)
	return nil
}

// GetFlow Flow Getter
func (r AlibabaaliqinflowwalletconsumeAPIRequest) GetFlow() int64 {
	return r._flow
}
