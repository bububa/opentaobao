package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaoumprangedelete 删除活动范围
// taobao.ump.range.delete
//
// 去指先前指定在某项活动中，某些类型的物品参加或者不参加活动的设置
func Taobaoumprangedelete(clt *core.SDKClient, req *promotion.TaobaoumprangedeleteAPIRequest, session string) (*promotion.TaobaoumprangedeleteAPIResponse, error) {
	var resp promotion.TaobaoumprangedeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
