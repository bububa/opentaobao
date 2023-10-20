package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Taobaologisticsinstanttracesearch 物流详情查询
// taobao.logistics.instant.trace.search
//
// 物流详情查询
func Taobaologisticsinstanttracesearch(clt *core.SDKClient, req *tblogistics.TaobaologisticsinstanttracesearchAPIRequest, session string) (*tblogistics.TaobaologisticsinstanttracesearchAPIResponse, error) {
	var resp tblogistics.TaobaologisticsinstanttracesearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
