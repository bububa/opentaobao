package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelbnbpromoget 民宿查询营销活动
// taobao.xhotel.bnbpromo.get
//
// 民宿查询营销活动
func Taobaoxhotelbnbpromoget(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelbnbpromogetAPIRequest, session string) (*xhotelitem.TaobaoxhotelbnbpromogetAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelbnbpromogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
