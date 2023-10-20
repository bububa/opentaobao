package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaoumpmbbslistget 通过ids列表获取营销积木块列表
// taobao.ump.mbbs.list.get
//
// 通过营销积木id列表来获取营销积木块列表。
func Taobaoumpmbbslistget(clt *core.SDKClient, req *promotion.TaobaoumpmbbslistgetAPIRequest, session string) (*promotion.TaobaoumpmbbslistgetAPIResponse, error) {
	var resp promotion.TaobaoumpmbbslistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
