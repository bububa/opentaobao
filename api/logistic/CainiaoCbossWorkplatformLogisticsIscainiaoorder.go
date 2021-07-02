package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// CainiaoCbossWorkplatformLogisticsIscainiaoorder 根据交易单号判断是否为菜鸟发货订单
// cainiao.cboss.workplatform.logistics.iscainiaoorder
//
// 根据交易单号判断是否为菜鸟发货订单
func CainiaoCbossWorkplatformLogisticsIscainiaoorder(clt *core.SDKClient, req *logistic.CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest, session string) (*logistic.CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse, error) {
	var resp logistic.CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
