package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelbnbhouseadd 民宿门店信息添加
// taobao.xhotel.bnbhouse.add
//
// 添加和更新民宿门店的信息
func Taobaoxhotelbnbhouseadd(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelbnbhouseaddAPIRequest, session string) (*xhotelitem.TaobaoxhotelbnbhouseaddAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelbnbhouseaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
