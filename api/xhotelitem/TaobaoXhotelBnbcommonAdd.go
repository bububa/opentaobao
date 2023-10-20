package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelbnbcommonadd 通用调用top接口
// taobao.xhotel.bnbcommon.add
//
// 通用调用top接口
func Taobaoxhotelbnbcommonadd(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelbnbcommonaddAPIRequest, session string) (*xhotelitem.TaobaoxhotelbnbcommonaddAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelbnbcommonaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
