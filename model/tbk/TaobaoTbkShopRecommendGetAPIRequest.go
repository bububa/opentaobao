package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkShopRecommendGetAPIRequest
淘宝客-公用-店铺关联推荐 API请求
taobao.tbk.shop.recommend.get

入参卖家id，可推荐与此店铺相关联的相关店铺。 */
type TaobaoTbkShopRecommendGetAPIRequest struct {
	model.Params
	// 需返回的字段列表
	_fields string
	// 卖家Id
	_userId int64
	// 返回数量，默认20，最大值40
	_count int64
	// 链接形式：1：PC，2：无线，默认：１
	_platform int64
}

// New
