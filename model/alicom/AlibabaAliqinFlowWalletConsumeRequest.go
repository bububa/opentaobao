package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
流量扣减 API请求
alibaba.aliqin.flow.wallet.consume

流量钱包流量扣减接口
*/
type AlibabaAliqinFlowWalletConsumeRequest struct {
    model.Params
    // 扣减流量值
    flow   int64
    // 扣减流水号
    serialNo   string
    // 扣减原因
    reason   string
    // 备注
    remark   string
}

// 初始化AlibabaAliqinFlowWalletConsumeRequest对象
func NewAlibabaAliqinFlowWalletConsumeRequest() *AlibabaAliqinFlowWalletConsumeRequest{
    return &AlibabaAliqinFlowWalletConsumeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowWalletConsumeRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.wallet.consume"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowWalletConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Flow Setter
// 扣减流量值
func (r *AlibabaAliqinFlowWalletConsumeRequest) SetFlow(flow int64) error {
    r.flow = flow
    r.Set("flow", flow)
    return nil
}

// Flow Getter
func (r AlibabaAliqinFlowWalletConsumeRequest) GetFlow() int64 {
    return r.flow
}
// SerialNo Setter
// 扣减流水号
func (r *AlibabaAliqinFlowWalletConsumeRequest) SetSerialNo(serialNo string) error {
    r.serialNo = serialNo
    r.Set("serial_no", serialNo)
    return nil
}

// SerialNo Getter
func (r AlibabaAliqinFlowWalletConsumeRequest) GetSerialNo() string {
    return r.serialNo
}
// Reason Setter
// 扣减原因
func (r *AlibabaAliqinFlowWalletConsumeRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

// Reason Getter
func (r AlibabaAliqinFlowWalletConsumeRequest) GetReason() string {
    return r.reason
}
// Remark Setter
// 备注
func (r *AlibabaAliqinFlowWalletConsumeRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

// Remark Getter
func (r AlibabaAliqinFlowWalletConsumeRequest) GetRemark() string {
    return r.remark
}
