package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrfulfilllogisticsquery 定时送和极速达配送物流信息查询
// tmall.nr.fulfill.logistics.query
//
// 发布定时送&amp;极速达物流信息查询服务
func Tmallnrfulfilllogisticsquery(clt *core.SDKClient, req *tmallnr.TmallnrfulfilllogisticsqueryAPIRequest, session string) (*tmallnr.TmallnrfulfilllogisticsqueryAPIResponse, error) {
	var resp tmallnr.TmallnrfulfilllogisticsqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
