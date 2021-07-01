package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOcOrderApUpdateAPIRequest
按OC订单分账 API请求
taobao.oc.order.ap.update

对OC订单执行分账操作 */
type TaobaoOcOrderApUpdateAPIRequest struct {
	model.Params
	// 调用创建OC订单接口生成的id
	_ocOrderId int64
}

// New
