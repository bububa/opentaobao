package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoNextoneLogisticsSignUpdateAPIRequest
AG物流签收状态写接口 API请求
taobao.nextone.logistics.sign.update

商家上传退货的签收状态给AG */
type TaobaoNextoneLogisticsSignUpdateAPIRequest struct {
	model.Params
	// 退款编号
	_refundId int64
	// 货物签收状态
	_signStatus int64
}

// New
