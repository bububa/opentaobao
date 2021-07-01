package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrtItemGetAPIRequest
家装新零售商品信息查询 API请求
tmall.nrt.item.get

查询新零售商品信息 */
type TmallNrtItemGetAPIRequest struct {
	model.Params
	// 城市站id
	_boothId int64
	// 商品id
	_itemId int64
}

// New
