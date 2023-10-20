package rail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rail"
)

// Alitriprailirserviceget 国际火车票仓位坐席查询
// alitrip.rail.ir.service.get
//
// 国际火车票标准仓位坐席查询
func Alitriprailirserviceget(clt *core.SDKClient, req *rail.AlitriprailirservicegetAPIRequest, session string) (*rail.AlitriprailirservicegetAPIResponse, error) {
	var resp rail.AlitriprailirservicegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
