package wdklogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdklogistics"
)

/* AlibabaWdkLogisticsPusPickupCararrived
自提业务-车辆到达上报车牌号
alibaba.wdk.logistics.pus.pickup.cararrived

自提业务-汽车自提,车辆到达上报车牌号 */
func AlibabaWdkLogisticsPusPickupCararrived(clt *core.SDKClient, req *wdklogistics.AlibabaWdkLogisticsPusPickupCararrivedAPIRequest, session string) (*wdklogistics.AlibabaWdkLogisticsPusPickupCararrivedAPIResponse, error) {
	var resp wdklogistics.AlibabaWdkLogisticsPusPickupCararrivedAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
