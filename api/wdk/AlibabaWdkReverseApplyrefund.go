package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkreverseapplyrefund 逆向申请接口
// alibaba.wdk.reverse.applyrefund
//
// 逆向渲染
func Alibabawdkreverseapplyrefund(clt *core.SDKClient, req *wdk.AlibabawdkreverseapplyrefundAPIRequest, session string) (*wdk.AlibabawdkreverseapplyrefundAPIResponse, error) {
	var resp wdk.AlibabawdkreverseapplyrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
