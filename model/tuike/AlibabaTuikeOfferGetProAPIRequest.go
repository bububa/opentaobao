package tuike

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTuikeOfferGetProAPIRequest
推广商品查询接口2.0 API请求
alibaba.tuike.offer.get.pro

查询1688推客平台卖家推广中的商品信息 */
type AlibabaTuikeOfferGetProAPIRequest struct {
	model.Params
	// 用户ID
	_loginId string
	// 标识调用方
	_isvCode string
	// 搜索查询参数(json)
	_queryString string
}

// New
