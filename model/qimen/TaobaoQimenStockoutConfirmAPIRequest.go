package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStockoutConfirmAPIRequest
出库单确认接口 API请求
taobao.qimen.stockout.confirm

货品出库后，WMS将状态回传给ERP */
type TaobaoQimenStockoutConfirmAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenStockoutConfirmStruct
}

// New
