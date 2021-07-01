package guoguo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest
兜底派送订单的运单号回传接口 API请求
cainiao.guoguo.backup.graborder.submitmailno

快递公司回传订单号和运单号给菜鸟裹裹 */
type CainiaoGuoguoBackupGraborderSubmitmailnoAPIRequest struct {
	model.Params
	// 菜鸟物流订单号
	_orderCode string
	// 运单号
	_mailNo string
}

// New
