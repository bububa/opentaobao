package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderGuideDataGetAPIRequest
获取全渠道导购产品数据 API请求
taobao.omniorder.guide.data.get

获取全渠道导购产品，目前包括随心购、随身购扫码、加购和交易数据。 */
type TaobaoOmniorderGuideDataGetAPIRequest struct {
	model.Params
	// detail_smg_scan: 扫码购扫码明细;detail_smg_cart: 扫码购加购明细;detail_smg_order: 扫码购订单明细;detail_sxg_search: 随心购搜索明细;detail_sxg_view_item: 随心购商品浏览明细;detail_sxg_cart: 随心购加购明细;detail_sxg_order: 随心购订单明细
	_type string
	// 拉取数据开始时间
	_startTime string
	// 页码，从1开始
	_pageNo int64
	// 每页数量，不能大于1000
	_pageSize int64
}

// New
