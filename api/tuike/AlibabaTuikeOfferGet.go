package tuike

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuike"
)

// AlibabaTuikeOfferGet 推广商品查询接口
// alibaba.tuike.offer.get
//
// 查询1688推客平台卖家推广中的商品信息
func AlibabaTuikeOfferGet(clt *core.SDKClient, req *tuike.AlibabaTuikeOfferGetAPIRequest, session string) (*tuike.AlibabaTuikeOfferGetAPIResponse, error) {
	var resp tuike.AlibabaTuikeOfferGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
