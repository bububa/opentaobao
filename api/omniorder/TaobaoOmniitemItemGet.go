package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniitemitemget 获取全渠道门店商品
// taobao.omniitem.item.get
//
// 通过门店id/类目id/商品id单个或多个参数组合查询全渠道门店商品信息
func Taobaoomniitemitemget(clt *core.SDKClient, req *omniorder.TaobaoomniitemitemgetAPIRequest, session string) (*omniorder.TaobaoomniitemitemgetAPIResponse, error) {
	var resp omniorder.TaobaoomniitemitemgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
