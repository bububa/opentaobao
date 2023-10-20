package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionmisccommonitemactivityupdate 修改通用单品优惠活动
// taobao.promotionmisc.common.item.activity.update
//
// 修改通用单品优惠活动。
// 1、该接口只修改活动基本信息，如需要增加、删除参与该活动的商品及优惠，请调用taobao.promotionmisc.common.item.detail.add和taobao.promotionmisc.common.item.detail.delete接口
// 2、使用该接口时需要把未做修改的字段值也传入
func Taobaopromotionmisccommonitemactivityupdate(clt *core.SDKClient, req *promotion.TaobaopromotionmisccommonitemactivityupdateAPIRequest, session string) (*promotion.TaobaopromotionmisccommonitemactivityupdateAPIResponse, error) {
	var resp promotion.TaobaopromotionmisccommonitemactivityupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
