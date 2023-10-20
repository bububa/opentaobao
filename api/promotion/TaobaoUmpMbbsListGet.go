package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpMbbsListGet 通过ids列表获取营销积木块列表
// taobao.ump.mbbs.list.get
//
// 通过营销积木id列表来获取营销积木块列表。
func TaobaoUmpMbbsListGet(clt *core.SDKClient, req *promotion.TaobaoUmpMbbsListGetAPIRequest, session string) (*promotion.TaobaoUmpMbbsListGetAPIResponse, error) {
	var resp promotion.TaobaoUmpMbbsListGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
