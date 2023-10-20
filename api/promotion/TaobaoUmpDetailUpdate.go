package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaoumpdetailupdate 修改活动详情
// taobao.ump.detail.update
//
// 更新活动详情
func Taobaoumpdetailupdate(clt *core.SDKClient, req *promotion.TaobaoumpdetailupdateAPIRequest, session string) (*promotion.TaobaoumpdetailupdateAPIResponse, error) {
	var resp promotion.TaobaoumpdetailupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
