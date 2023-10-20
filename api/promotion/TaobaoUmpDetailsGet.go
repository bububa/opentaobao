package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaoumpdetailsget 查询活动详情列表
// taobao.ump.details.get
//
// 分页查询优惠详情列表
func Taobaoumpdetailsget(clt *core.SDKClient, req *promotion.TaobaoumpdetailsgetAPIRequest, session string) (*promotion.TaobaoumpdetailsgetAPIResponse, error) {
	var resp promotion.TaobaoumpdetailsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
