package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaowlbsubscriptionquery 查询商家定购的所有服务
// taobao.wlb.subscription.query
//
// 查询商家定购的所有服务,可通过入参状态来筛选
func Taobaowlbsubscriptionquery(clt *core.SDKClient, req *wlb.TaobaowlbsubscriptionqueryAPIRequest, session string) (*wlb.TaobaowlbsubscriptionqueryAPIResponse, error) {
	var resp wlb.TaobaowlbsubscriptionqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
