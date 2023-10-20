package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// CainiaoCbossWorkplatformLogisticsIscainiaoorder 根据交易单号判断是否为菜鸟发货订单
// cainiao.cboss.workplatform.logistics.iscainiaoorder
//
// 根据交易单号判断是否为菜鸟发货订单
func CainiaoCbossWorkplatformLogisticsIscainiaoorder(clt *core.SDKClient, req *logistic.CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest, resp *logistic.CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
