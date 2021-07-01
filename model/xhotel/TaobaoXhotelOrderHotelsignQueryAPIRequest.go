package xhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderHotelsignQueryAPIRequest
获取直连酒店（客栈）签约上线进度信息 API请求
taobao.xhotel.order.hotelsign.query

获取直连酒店（客栈）签约上线进度信息 */
type TaobaoXhotelOrderHotelsignQueryAPIRequest struct {
	model.Params
	// 请求流水
	_outUuid string
	// 商家酒店编码
	_hotelCode string
	// 商家vendor
	_vendor string
	// 1
	_type string
	// 页码
	_pageNo int64
}

// New
