package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkumsoutboundprocessget 出库业务UMS异步处理结果返回
// alibaba.wdk.ums.outbound.process.get
//
// 出库业务UMS异步处理结果返回
func Alibabawdkumsoutboundprocessget(clt *core.SDKClient, req *wdk.AlibabawdkumsoutboundprocessgetAPIRequest, session string) (*wdk.AlibabawdkumsoutboundprocessgetAPIResponse, error) {
	var resp wdk.AlibabawdkumsoutboundprocessgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
