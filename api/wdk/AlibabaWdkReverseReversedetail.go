package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkreversereversedetail 退款详情
// alibaba.wdk.reverse.reversedetail
//
// 退款详情
func Alibabawdkreversereversedetail(clt *core.SDKClient, req *wdk.AlibabawdkreversereversedetailAPIRequest, session string) (*wdk.AlibabawdkreversereversedetailAPIResponse, error) {
	var resp wdk.AlibabawdkreversereversedetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
