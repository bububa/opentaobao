package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionmisccommonitemdetaildelete 删除通用单品优惠详情
// taobao.promotionmisc.common.item.detail.delete
//
// 删除通用单品优惠详情。
func Taobaopromotionmisccommonitemdetaildelete(clt *core.SDKClient, req *promotion.TaobaopromotionmisccommonitemdetaildeleteAPIRequest, session string) (*promotion.TaobaopromotionmisccommonitemdetaildeleteAPIResponse, error) {
	var resp promotion.TaobaopromotionmisccommonitemdetaildeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
