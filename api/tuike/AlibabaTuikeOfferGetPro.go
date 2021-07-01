package tuike

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuike"
)

/* AlibabaTuikeOfferGetPro
推广商品查询接口2.0
alibaba.tuike.offer.get.pro

查询1688推客平台卖家推广中的商品信息 */
func AlibabaTuikeOfferGetPro(clt *core.SDKClient, req *tuike.AlibabaTuikeOfferGetProAPIRequest, session string) (*tuike.AlibabaTuikeOfferGetProAPIResponse, error) {
	var resp tuike.AlibabaTuikeOfferGetProAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
