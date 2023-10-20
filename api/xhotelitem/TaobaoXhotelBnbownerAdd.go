package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelbnbowneradd 民宿房东信息添加
// taobao.xhotel.bnbowner.add
//
// 添加和更新民宿房东的信息
func Taobaoxhotelbnbowneradd(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelbnbowneraddAPIRequest, session string) (*xhotelitem.TaobaoxhotelbnbowneraddAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelbnbowneraddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
