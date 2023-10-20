package tuike

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuike"
)

// AlibabaTuikeOfferGet 推广商品查询接口
// alibaba.tuike.offer.get
//
// 查询1688推客平台卖家推广中的商品信息
func AlibabaTuikeOfferGet(clt *core.SDKClient, req *tuike.AlibabaTuikeOfferGetAPIRequest, resp *tuike.AlibabaTuikeOfferGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
