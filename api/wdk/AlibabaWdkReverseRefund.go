package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkreverserefund 退款打款
// alibaba.wdk.reverse.refund
//
// 五道口退款打款开放能力接口
func Alibabawdkreverserefund(clt *core.SDKClient, req *wdk.AlibabawdkreverserefundAPIRequest, session string) (*wdk.AlibabawdkreverserefundAPIResponse, error) {
	var resp wdk.AlibabawdkreverserefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
