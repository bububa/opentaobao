package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest
商家数据-选品整体概况 API请求
taobao.xhotel.item.selection.seller.stat.summary

商家数据-选品整体概况 */
type TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest struct {
	model.Params
	// vendor 默认为all
	_vendor string
	// 日期  默认为昨天
	_date string
	// hid  默认为all
	_hid string
	// supplier 默认为all
	_supplier string
	// 外部酒店编码
	_outHid string
}

// New
