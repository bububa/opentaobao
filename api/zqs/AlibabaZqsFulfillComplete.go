package zqs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/zqs"
)

// Alibabazqsfulfillcomplete 周期购履约完成接口
// alibaba.zqs.fulfill.complete
//
// 周期购履约完成接口
func Alibabazqsfulfillcomplete(clt *core.SDKClient, req *zqs.AlibabazqsfulfillcompleteAPIRequest, session string) (*zqs.AlibabazqsfulfillcompleteAPIResponse, error) {
	var resp zqs.AlibabazqsfulfillcompleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
