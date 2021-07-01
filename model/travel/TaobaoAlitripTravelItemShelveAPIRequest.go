package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelItemShelveAPIRequest
【API3.0】度假线路商品上下架接口 API请求
taobao.alitrip.travel.item.shelve

旅行度假新商品发布接口 第三版：度假商品上下架接口
注意：定时上下架功能，目前只支持接送、租车类目。 */
type TaobaoAlitripTravelItemShelveAPIRequest struct {
	model.Params
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
	// 1-上架 0-下架
	_itemStatus int64
	// 指定定时上架时间，格式：yyyy-MM-dd HH:mm:ss。若不设置该值且item_status为1，则表示立即上架。
	_onlineTime string
}

// New
