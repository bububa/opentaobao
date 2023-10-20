package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaoumpactivitiesget 查询活动列表
// taobao.ump.activities.get
//
// 查询活动列表
func Taobaoumpactivitiesget(clt *core.SDKClient, req *promotion.TaobaoumpactivitiesgetAPIRequest, session string) (*promotion.TaobaoumpactivitiesgetAPIResponse, error) {
	var resp promotion.TaobaoumpactivitiesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
