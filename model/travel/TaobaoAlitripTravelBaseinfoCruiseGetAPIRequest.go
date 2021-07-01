package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest
【API3.0】度假线路商品发布时基础信息获取接口：邮轮扩展信息获取 API请求
taobao.alitrip.travel.baseinfo.cruise.get

旅行度假新商品发布时可用的扩展接口，用于获取邮轮类目相关扩展信息。 */
type TaobaoAlitripTravelBaseinfoCruiseGetAPIRequest struct {
	model.Params
	// true-获取国际邮轮类目扩展信息；false-获取国内邮轮类目扩展信息
	_isOverseas bool
}

// New
