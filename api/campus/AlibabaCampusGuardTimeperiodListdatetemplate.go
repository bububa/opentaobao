package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusguardtimeperiodlistdatetemplate 门禁控制器查询日期模版
// alibaba.campus.guard.timeperiod.listdatetemplate
//
// 门禁控制器查询日期模版
func Alibabacampusguardtimeperiodlistdatetemplate(clt *core.SDKClient, req *campus.AlibabacampusguardtimeperiodlistdatetemplateAPIRequest, session string) (*campus.AlibabacampusguardtimeperiodlistdatetemplateAPIResponse, error) {
	var resp campus.AlibabacampusguardtimeperiodlistdatetemplateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
