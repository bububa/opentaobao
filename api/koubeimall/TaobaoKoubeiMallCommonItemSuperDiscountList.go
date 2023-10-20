package koubeimall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// Taobaokoubeimallcommonitemsuperdiscountlist 查询商圈内的超值特惠商品信息
// taobao.koubei.mall.common.item.super.discount.list
//
// 查询商圈超值特惠商品信息列表
func Taobaokoubeimallcommonitemsuperdiscountlist(clt *core.SDKClient, req *koubeimall.TaobaokoubeimallcommonitemsuperdiscountlistAPIRequest, session string) (*koubeimall.TaobaokoubeimallcommonitemsuperdiscountlistAPIResponse, error) {
	var resp koubeimall.TaobaokoubeimallcommonitemsuperdiscountlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
