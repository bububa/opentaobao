package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrFulfillLogisticsQuery 定时送和极速达配送物流信息查询
// tmall.nr.fulfill.logistics.query
//
// 发布定时送&amp;极速达物流信息查询服务
func TmallNrFulfillLogisticsQuery(clt *core.SDKClient, req *tmallnr.TmallNrFulfillLogisticsQueryAPIRequest, session string) (*tmallnr.TmallNrFulfillLogisticsQueryAPIResponse, error) {
	var resp tmallnr.TmallNrFulfillLogisticsQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
