package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJdsHluserUpdateAPIRequest
订单全链路用户信息修改 API请求
taobao.jds.hluser.update

订单全链路用户信息修改，比如是否开放买家端展示 */
type TaobaoJdsHluserUpdateAPIRequest struct {
	model.Params
	// 回流信息是否开通买家端展示,可选值open,close
	_openForBuyer string
	// 为空,则默认是X_TO_SYSTEM,X_WAIT_ALLOCATION,X_OUT_WAREHOUSE 可选值是X_DOWNLOADED,X_TO_SYSTEM,X_SERVICE_AUDITED,X_FINANCE_AUDITED,X_ALLOCATION_NOTIFIED,X_WAIT_ALLOCATION,X_SORT_PRINTED,X_SEND_PRINTED,X_LOGISTICS_PRINTED,X_SORTED,X_EXAMINED,X_PACKAGED,X_WEIGHED,X_OUT_WAREHOUS
	_openNodes string
}

// New
