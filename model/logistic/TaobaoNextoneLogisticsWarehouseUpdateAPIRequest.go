package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoNextoneLogisticsWarehouseUpdateAPIRequest
AG退货入仓状态写接口 API请求
taobao.nextone.logistics.warehouse.update

商家上传退货入仓状态给ag */
type TaobaoNextoneLogisticsWarehouseUpdateAPIRequest struct {
	model.Params
	// 退款编号
	_refundId int64
	// 退货入仓状态 1.已入仓
	_warehouseStatus int64
}

// New
