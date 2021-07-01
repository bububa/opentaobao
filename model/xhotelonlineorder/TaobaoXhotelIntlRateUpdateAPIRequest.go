package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelIntlRateUpdateAPIRequest
不落库商家推送更新酒店rate API请求
taobao.xhotel.intl.rate.update

商家主动推送不落库商品的酒店信息 */
type TaobaoXhotelIntlRateUpdateAPIRequest struct {
	model.Params
	// rate更新参数
	_updateRateParam *UpdateRateParam
}

// New
