package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaoumpactivityupdate 修改活动信息
// taobao.ump.activity.update
//
// 修改营销活动
func Taobaoumpactivityupdate(clt *core.SDKClient, req *promotion.TaobaoumpactivityupdateAPIRequest, session string) (*promotion.TaobaoumpactivityupdateAPIResponse, error) {
	var resp promotion.TaobaoumpactivityupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
