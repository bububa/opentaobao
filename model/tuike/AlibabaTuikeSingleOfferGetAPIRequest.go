package tuike

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTuikeSingleOfferGetAPIRequest
查询1688推客平台卖家推广中的商品信息 API请求
alibaba.tuike.single.offer.get

查询单个推客商品信息的接口 */
type AlibabaTuikeSingleOfferGetAPIRequest struct {
	model.Params
	// 推客id
	_loginId string
	// 商品id
	_offerId int64
}

// New
