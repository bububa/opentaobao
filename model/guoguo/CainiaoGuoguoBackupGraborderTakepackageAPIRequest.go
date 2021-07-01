package guoguo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGuoguoBackupGraborderTakepackageAPIRequest
兜底派送订单的揽件接口 API请求
cainiao.guoguo.backup.graborder.takepackage

快递公司回传订单号和四位取件码给菜鸟裹裹 */
type CainiaoGuoguoBackupGraborderTakepackageAPIRequest struct {
	model.Params
	// 物流订单号
	_orderCode string
	// 包裹四位码
	_packageCode string
}

// New
