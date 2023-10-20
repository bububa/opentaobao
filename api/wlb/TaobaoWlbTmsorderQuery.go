package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaowlbtmsorderquery 通过物流订单编号查询物流信息
// taobao.wlb.tmsorder.query
//
// 通过物流订单编号分页查询物流信息
func Taobaowlbtmsorderquery(clt *core.SDKClient, req *wlb.TaobaowlbtmsorderqueryAPIRequest, session string) (*wlb.TaobaowlbtmsorderqueryAPIResponse, error) {
	var resp wlb.TaobaowlbtmsorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
