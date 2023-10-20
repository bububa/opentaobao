package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TmallPromotagTaguserJudge 用户标签判断接口
// tmall.promotag.taguser.judge
//
// 查询用户是否有标签
func TmallPromotagTaguserJudge(clt *core.SDKClient, req *promotion.TmallPromotagTaguserJudgeAPIRequest, resp *promotion.TmallPromotagTaguserJudgeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
