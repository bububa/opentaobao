package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
盒马帮发货信息回传接口 API请求
alibaba.wdk.wholesale.outboundorder.commit

盒马帮发货信息回传接口
*/
type AlibabaWdkWholesaleOutboundorderCommitRequest struct {
    model.Params
    // 发货信息参数
    outboundInfoCommitReq   *OutboundInfoCommitReq
}

// 初始化AlibabaWdkWholesaleOutboundorderCommitRequest对象
func NewAlibabaWdkWholesaleOutboundorderCommitRequest() *AlibabaWdkWholesaleOutboundorderCommitRequest{
    return &AlibabaWdkWholesaleOutboundorderCommitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkWholesaleOutboundorderCommitRequest) GetApiMethodName() string {
    return "alibaba.wdk.wholesale.outboundorder.commit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkWholesaleOutboundorderCommitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutboundInfoCommitReq Setter
// 发货信息参数
func (r *AlibabaWdkWholesaleOutboundorderCommitRequest) SetOutboundInfoCommitReq(outboundInfoCommitReq *OutboundInfoCommitReq) error {
    r.outboundInfoCommitReq = outboundInfoCommitReq
    r.Set("outbound_info_commit_req", outboundInfoCommitReq)
    return nil
}

// OutboundInfoCommitReq Getter
func (r AlibabaWdkWholesaleOutboundorderCommitRequest) GetOutboundInfoCommitReq() *OutboundInfoCommitReq {
    return r.outboundInfoCommitReq
}
