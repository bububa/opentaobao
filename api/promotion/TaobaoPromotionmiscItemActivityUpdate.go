package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
修改无条件单品优惠活动 
taobao.promotionmisc.item.activity.update

修改无条件单品优惠活动。<br/>1、该接口只修改活动基本信息和打折信息，如需要增加、删除参与该活动的商品，请调用taobao.promotionmisc.activity.range.add和taobao.promotionmisc.activity.range.remove接口。 <br/>2、使用该接口时需要同时把未做修改的字段值也传入。 <br/><br/>3、该接口受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。
*/
func TaobaoPromotionmiscItemActivityUpdate(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscItemActivityUpdateRequest, session string) (*promotion.TaobaoPromotionmiscItemActivityUpdateResponse, error) {
    var resp promotion.TaobaoPromotionmiscItemActivityUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
