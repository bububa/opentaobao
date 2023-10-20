package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkbmtradeactivityquery 品牌营销的订单活动信息查询
// alibaba.wdk.bm.trade.activity.query
//
// 品牌营销的订单活动信息查询
func Alibabawdkbmtradeactivityquery(clt *core.SDKClient, req *wdk.AlibabawdkbmtradeactivityqueryAPIRequest, session string) (*wdk.AlibabawdkbmtradeactivityqueryAPIResponse, error) {
	var resp wdk.AlibabawdkbmtradeactivityqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
