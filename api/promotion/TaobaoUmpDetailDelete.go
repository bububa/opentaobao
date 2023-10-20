package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaoumpdetaildelete 删除活动详情
// taobao.ump.detail.delete
//
// 删除活动详情
func Taobaoumpdetaildelete(clt *core.SDKClient, req *promotion.TaobaoumpdetaildeleteAPIRequest, session string) (*promotion.TaobaoumpdetaildeleteAPIResponse, error) {
	var resp promotion.TaobaoumpdetaildeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
