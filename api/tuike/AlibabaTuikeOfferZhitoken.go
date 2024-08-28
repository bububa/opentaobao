package tuike

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuike"
)

// AlibabaTuikeOfferZhitoken 生成阿里口令
// alibaba.tuike.offer.zhitoken
//
// 推荐链接生产吱口令
func AlibabaTuikeOfferZhitoken(ctx context.Context, clt *core.SDKClient, req *tuike.AlibabaTuikeOfferZhitokenAPIRequest, resp *tuike.AlibabaTuikeOfferZhitokenAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
