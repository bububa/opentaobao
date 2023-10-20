package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkreverseloadfeatchorder 取货详情
// alibaba.wdk.reverse.load.featchorder
//
// 取货详情
func Alibabawdkreverseloadfeatchorder(clt *core.SDKClient, req *wdk.AlibabawdkreverseloadfeatchorderAPIRequest, session string) (*wdk.AlibabawdkreverseloadfeatchorderAPIResponse, error) {
	var resp wdk.AlibabawdkreverseloadfeatchorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
