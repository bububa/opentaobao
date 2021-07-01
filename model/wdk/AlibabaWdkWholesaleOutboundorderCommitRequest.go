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
type AlibabaWdkWholesaleOutboundorderCommitAPIRequest struct {
    model.Params
    // 发货信息参数
    _outboundInfoCommitReq   *OutboundInfoCommitReq
}

// 初始化AlibabaWdkWholesaleOutboundorderCommitAPIRequest对象
func NewAlibabaWdkWholesaleOutboundorderCommitRequest() *AlibabaWdkWholesaleOutboundorderCommitAPIRequest{
    return &AlibabaWdkWholesaleOutboundorderCommitAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkWholesaleOutboundorderCommitAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.wholesale.outboundorder.commit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkWholesaleOutboundorderCommitAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutboundInfoCommitReq Setter
// 发货信息参数
func (r *AlibabaWdkWholesaleOutboundorderCommitAPIRequest) SetOutboundInfoCommitReq(_outboundInfoCommitReq *OutboundInfoCommitReq) error {
    r._outboundInfoCommitReq = _outboundInfoCommitReq
    r.Set("outbound_info_commit_req", _outboundInfoCommitReq)
    return nil
}

// OutboundInfoCommitReq Getter
func (r AlibabaWdkWholesaleOutboundorderCommitAPIRequest) GetOutboundInfoCommitReq() *OutboundInfoCommitReq {
    return r._outboundInfoCommitReq
}
