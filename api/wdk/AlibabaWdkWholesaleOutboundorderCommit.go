package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkWholesaleOutboundorderCommit 盒马帮发货信息回传接口
// alibaba.wdk.wholesale.outboundorder.commit
//
// 盒马帮发货信息回传接口
func AlibabaWdkWholesaleOutboundorderCommit(clt *core.SDKClient, req *wdk.AlibabaWdkWholesaleOutboundorderCommitAPIRequest, session string) (*wdk.AlibabaWdkWholesaleOutboundorderCommitAPIResponse, error) {
	var resp wdk.AlibabaWdkWholesaleOutboundorderCommitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
