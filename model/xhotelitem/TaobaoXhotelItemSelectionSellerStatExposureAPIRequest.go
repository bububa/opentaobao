package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelItemSelectionSellerStatExposureAPIRequest
选品曝光趋势 API请求
taobao.xhotel.item.selection.seller.stat.exposure

用于提供给商家获取选品曝光趋势 */
type TaobaoXhotelItemSelectionSellerStatExposureAPIRequest struct {
	model.Params
	// 日期 默认为昨天
	_date string
	// hid  默认为all
	_hid string
	// 默认为all
	_vendor string
	// 默认为all
	_supplier string
	// 酒店编码
	_outHid string
}

// New
