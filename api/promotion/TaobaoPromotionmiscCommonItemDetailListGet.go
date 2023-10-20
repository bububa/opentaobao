package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionmisccommonitemdetaillistget 查询通用单品优惠详情列表
// taobao.promotionmisc.common.item.detail.list.get
//
// 查询通用单品优惠详情列表。
func Taobaopromotionmisccommonitemdetaillistget(clt *core.SDKClient, req *promotion.TaobaopromotionmisccommonitemdetaillistgetAPIRequest, session string) (*promotion.TaobaopromotionmisccommonitemdetaillistgetAPIResponse, error) {
	var resp promotion.TaobaopromotionmisccommonitemdetaillistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
