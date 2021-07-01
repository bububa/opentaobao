package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest
根据交易单号判断是否为菜鸟发货订单 API请求
cainiao.cboss.workplatform.logistics.iscainiaoorder

根据交易单号判断是否为菜鸟发货订单 */
type CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIRequest struct {
	model.Params
	// 交易单号
	_tradeId string
}

// New
