package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusguardtimeperiodlisttimerule 门禁控制器查询时间规则
// alibaba.campus.guard.timeperiod.listtimerule
//
// 门禁控制器查询时间规则
func Alibabacampusguardtimeperiodlisttimerule(clt *core.SDKClient, req *campus.AlibabacampusguardtimeperiodlisttimeruleAPIRequest, session string) (*campus.AlibabacampusguardtimeperiodlisttimeruleAPIResponse, error) {
	var resp campus.AlibabacampusguardtimeperiodlisttimeruleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
