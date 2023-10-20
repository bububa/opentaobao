package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Alibabalsycrmactivitysellerinfo 商家信息推送
// alibaba.lsy.crm.activity.sellerinfo
//
// 本地团商家信息推送
func Alibabalsycrmactivitysellerinfo(clt *core.SDKClient, req *tmallnr.AlibabalsycrmactivitysellerinfoAPIRequest, session string) (*tmallnr.AlibabalsycrmactivitysellerinfoAPIResponse, error) {
	var resp tmallnr.AlibabalsycrmactivitysellerinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
