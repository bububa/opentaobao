package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkwholesaleinboundordercommit 盒马帮退货信息回传接口
// alibaba.wdk.wholesale.inboundorder.commit
//
// 盒马帮退货信息回传接口
func Alibabawdkwholesaleinboundordercommit(clt *core.SDKClient, req *wdk.AlibabawdkwholesaleinboundordercommitAPIRequest, session string) (*wdk.AlibabawdkwholesaleinboundordercommitAPIResponse, error) {
	var resp wdk.AlibabawdkwholesaleinboundordercommitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
