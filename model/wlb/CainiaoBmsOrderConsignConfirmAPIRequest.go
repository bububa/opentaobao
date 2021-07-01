package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoBmsOrderConsignConfirmAPIRequest
BMS出库通知 API请求
cainiao.bms.order.consign.confirm

BMS出库后，通知ISV */
type CainiaoBmsOrderConsignConfirmAPIRequest struct {
	model.Params
	// 通知消息主体
	_content *BmsConsignOrderConfirm
}

// New
