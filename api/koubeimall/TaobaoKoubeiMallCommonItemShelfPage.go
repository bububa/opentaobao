package koubeimall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// Taobaokoubeimallcommonitemshelfpage 门店货架商品列表信息查询
// taobao.koubei.mall.common.item.shelf.page
//
// 查询口碑综合体内门店货架商品列表信息接口
func Taobaokoubeimallcommonitemshelfpage(clt *core.SDKClient, req *koubeimall.TaobaokoubeimallcommonitemshelfpageAPIRequest, session string) (*koubeimall.TaobaokoubeimallcommonitemshelfpageAPIResponse, error) {
	var resp koubeimall.TaobaokoubeimallcommonitemshelfpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
