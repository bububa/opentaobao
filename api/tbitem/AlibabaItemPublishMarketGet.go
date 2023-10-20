package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Alibabaitempublishmarketget 获取商家可发布商品的市场信息
// alibaba.item.publish.market.get
//
// 获取商家可发布商品的市场信息
func Alibabaitempublishmarketget(clt *core.SDKClient, req *tbitem.AlibabaitempublishmarketgetAPIRequest, session string) (*tbitem.AlibabaitempublishmarketgetAPIResponse, error) {
	var resp tbitem.AlibabaitempublishmarketgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
