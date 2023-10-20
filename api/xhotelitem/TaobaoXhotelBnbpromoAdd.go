package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelbnbpromoadd 自促活动申请接口
// taobao.xhotel.bnbpromo.add
//
// 自促活动申请接口
func Taobaoxhotelbnbpromoadd(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelbnbpromoaddAPIRequest, session string) (*xhotelitem.TaobaoxhotelbnbpromoaddAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelbnbpromoaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
