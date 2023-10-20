package wdklogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdklogistics"
)

// Alibabawdklogisticspuspickupcararrived 自提业务-车辆到达上报车牌号
// alibaba.wdk.logistics.pus.pickup.cararrived
//
// 自提业务-汽车自提,车辆到达上报车牌号
func Alibabawdklogisticspuspickupcararrived(clt *core.SDKClient, req *wdklogistics.AlibabawdklogisticspuspickupcararrivedAPIRequest, session string) (*wdklogistics.AlibabawdklogisticspuspickupcararrivedAPIResponse, error) {
	var resp wdklogistics.AlibabawdklogisticspuspickupcararrivedAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
