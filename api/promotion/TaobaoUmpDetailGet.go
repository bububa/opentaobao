package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaoumpdetailget 查询活动详情
// taobao.ump.detail.get
//
// 查询活动详情
func Taobaoumpdetailget(clt *core.SDKClient, req *promotion.TaobaoumpdetailgetAPIRequest, session string) (*promotion.TaobaoumpdetailgetAPIResponse, error) {
	var resp promotion.TaobaoumpdetailgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
