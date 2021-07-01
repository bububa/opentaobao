package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOcOrderCreateAPIRequest
创建OC订单 API请求
taobao.oc.order.create

创建OC订单接口 */
type TaobaoOcOrderCreateAPIRequest struct {
	model.Params
	// OC订单
	_paramOCOrder *OcOrder
}

// New
