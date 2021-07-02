package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsAddressReachable 判定服务是否可达
// taobao.logistics.address.reachable
//
// 根据输入的目标地址，判断服务是否可达。
// 现已支持筛单的快递公司共15家：中国邮政、EMS、国通、汇通、快捷、全峰、优速、圆通、宅急送、中通、顺丰、天天、韵达、德邦快递、申通
func TaobaoLogisticsAddressReachable(clt *core.SDKClient, req *logistic.TaobaoLogisticsAddressReachableAPIRequest, session string) (*logistic.TaobaoLogisticsAddressReachableAPIResponse, error) {
	var resp logistic.TaobaoLogisticsAddressReachableAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
