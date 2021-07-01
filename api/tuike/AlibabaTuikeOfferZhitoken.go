package tuike

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuike"
)

/* AlibabaTuikeOfferZhitoken
生成阿里口令
alibaba.tuike.offer.zhitoken

推荐链接生产吱口令 */
func AlibabaTuikeOfferZhitoken(clt *core.SDKClient, req *tuike.AlibabaTuikeOfferZhitokenAPIRequest, session string) (*tuike.AlibabaTuikeOfferZhitokenAPIResponse, error) {
	var resp tuike.AlibabaTuikeOfferZhitokenAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
