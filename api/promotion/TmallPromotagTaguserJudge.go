package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
用户标签判断接口 
tmall.promotag.taguser.judge

查询用户是否有标签
*/
func TmallPromotagTaguserJudge(clt *core.SDKClient, req *promotion.TmallPromotagTaguserJudgeRequest, session string) (*promotion.TmallPromotagTaguserJudgeResponse, error) {
    var resp promotion.TmallPromotagTaguserJudgeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
