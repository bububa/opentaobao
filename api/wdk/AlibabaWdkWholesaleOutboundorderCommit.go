package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkwholesaleoutboundordercommit 盒马帮发货信息回传接口
// alibaba.wdk.wholesale.outboundorder.commit
//
// 盒马帮发货信息回传接口
func Alibabawdkwholesaleoutboundordercommit(clt *core.SDKClient, req *wdk.AlibabawdkwholesaleoutboundordercommitAPIRequest, session string) (*wdk.AlibabawdkwholesaleoutboundordercommitAPIResponse, error) {
	var resp wdk.AlibabawdkwholesaleoutboundordercommitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
