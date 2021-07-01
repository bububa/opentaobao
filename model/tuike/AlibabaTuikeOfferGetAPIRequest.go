package tuike

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTuikeOfferGetAPIRequest
推广商品查询接口 API请求
alibaba.tuike.offer.get

查询1688推客平台卖家推广中的商品信息 */
type AlibabaTuikeOfferGetAPIRequest struct {
	model.Params
	// 标识调用方
	_isvCode string
	// 搜索查询参数(json)
	_queryString string
}

// New
