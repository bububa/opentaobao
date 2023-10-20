package tuike

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuike"
)

// AlibabaTuikeOfferZhitoken 生成阿里口令
// alibaba.tuike.offer.zhitoken
//
// 推荐链接生产吱口令
func AlibabaTuikeOfferZhitoken(clt *core.SDKClient, req *tuike.AlibabaTuikeOfferZhitokenAPIRequest, resp *tuike.AlibabaTuikeOfferZhitokenAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
