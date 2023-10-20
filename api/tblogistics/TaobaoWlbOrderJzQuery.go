package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaowlborderjzquery 家装业务查询物流公司api
// taobao.wlb.order.jz.query
//
// 家装业务查询物流公司api
func Taobaowlborderjzquery(clt *core.SDKClient, req *tblogistics.TaobaowlborderjzqueryAPIRequest, session string) (*tblogistics.TaobaowlborderjzqueryAPIResponse, error) {
	var resp tblogistics.TaobaowlborderjzqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
