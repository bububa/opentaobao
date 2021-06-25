package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
流量扣减 APIRequest
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

func NewAlibabaAliqinFlowWalletConsumeRequest() *AlibabaAliqinFlowWalletConsumeRequest{
    return &AlibabaAliqinFlowWalletConsumeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFlowWalletConsumeRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.wallet.consume"
}

func (r AlibabaAliqinFlowWalletConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFlowWalletConsumeRequest) SetFlow(flow int64) error {
    r.flow = flow
    r.Set("flow", flow)
    return nil
}

func (r AlibabaAliqinFlowWalletConsumeRequest) GetFlow() int64 {
    return r.flow
}

func (r *AlibabaAliqinFlowWalletConsumeRequest) SetSerialNo(serialNo string) error {
    r.serialNo = serialNo
    r.Set("serial_no", serialNo)
    return nil
}

func (r AlibabaAliqinFlowWalletConsumeRequest) GetSerialNo() string {
    return r.serialNo
}

func (r *AlibabaAliqinFlowWalletConsumeRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r AlibabaAliqinFlowWalletConsumeRequest) GetReason() string {
    return r.reason
}

func (r *AlibabaAliqinFlowWalletConsumeRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r AlibabaAliqinFlowWalletConsumeRequest) GetRemark() string {
    return r.remark
}

