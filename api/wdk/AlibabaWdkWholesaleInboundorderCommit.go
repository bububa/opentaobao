package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkWholesaleInboundorderCommit 盒马帮退货信息回传接口
// alibaba.wdk.wholesale.inboundorder.commit
//
// 盒马帮退货信息回传接口
func AlibabaWdkWholesaleInboundorderCommit(clt *core.SDKClient, req *wdk.AlibabaWdkWholesaleInboundorderCommitAPIRequest, session string) (*wdk.AlibabaWdkWholesaleInboundorderCommitAPIResponse, error) {
	var resp wdk.AlibabaWdkWholesaleInboundorderCommitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
