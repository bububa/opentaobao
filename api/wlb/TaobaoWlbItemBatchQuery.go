package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaowlbitembatchquery 批次库存查询接口
// taobao.wlb.item.batch.query
//
// 根据用户id，item id list和store code来查询商品库存信息和批次信息
func Taobaowlbitembatchquery(clt *core.SDKClient, req *wlb.TaobaowlbitembatchqueryAPIRequest, session string) (*wlb.TaobaowlbitembatchqueryAPIResponse, error) {
	var resp wlb.TaobaowlbitembatchqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
