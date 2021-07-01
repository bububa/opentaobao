package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest
供应链选品热销标准酒店覆盖情况 API请求
taobao.xhotel.item.selection.seller.stat.hotshid

供应链选品热销标准酒店覆盖情况 */
type TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest struct {
	model.Params
	// 日期  默认为昨天
	_date string
	// 酒店id  默认all
	_hid string
	// vendor  默认all
	_vendor string
	// supplier  默认all
	_supplier string
	// 酒店编码
	_outHid string
}

// New
