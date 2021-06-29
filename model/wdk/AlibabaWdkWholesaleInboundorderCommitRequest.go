package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
盒马帮退货信息回传接口 API请求
alibaba.wdk.wholesale.inboundorder.commit

盒马帮退货信息回传接口
*/
type AlibabaWdkWholesaleInboundorderCommitRequest struct {
    model.Params
    // 退货信息参数
    inboundInfoCommitReq   *InboundInfoCommitReq
}

// 初始化AlibabaWdkWholesaleInboundorderCommitRequest对象
func NewAlibabaWdkWholesaleInboundorderCommitRequest() *AlibabaWdkWholesaleInboundorderCommitRequest{
    return &AlibabaWdkWholesaleInboundorderCommitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkWholesaleInboundorderCommitRequest) GetApiMethodName() string {
    return "alibaba.wdk.wholesale.inboundorder.commit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkWholesaleInboundorderCommitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InboundInfoCommitReq Setter
// 退货信息参数
func (r *AlibabaWdkWholesaleInboundorderCommitRequest) SetInboundInfoCommitReq(inboundInfoCommitReq *InboundInfoCommitReq) error {
    r.inboundInfoCommitReq = inboundInfoCommitReq
    r.Set("inbound_info_commit_req", inboundInfoCommitReq)
    return nil
}

// InboundInfoCommitReq Getter
func (r AlibabaWdkWholesaleInboundorderCommitRequest) GetInboundInfoCommitReq() *InboundInfoCommitReq {
    return r.inboundInfoCommitReq
}
