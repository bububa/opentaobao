package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsceslIteminfoBatchInsertAPIRequest
按商家批量写入商品接口 API请求
taobao.uscesl.iteminfo.batch.insert

【电子价签】支持按照商家-门店维度批量写入商品数据 */
type TaobaoUsceslIteminfoBatchInsertAPIRequest struct {
	model.Params
	// 商品信息集合，限制500
	_itemList []ItemChangeBo
	// 门店ID
	_storeId int64
	// 商家KEY
	_bizBrandKey string
}

// New
