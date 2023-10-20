package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelbnbpromodelete 民宿卖家活动删除
// taobao.xhotel.bnbpromo.delete
//
// 民宿删除营销活动
func Taobaoxhotelbnbpromodelete(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelbnbpromodeleteAPIRequest, session string) (*xhotelitem.TaobaoxhotelbnbpromodeleteAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelbnbpromodeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
