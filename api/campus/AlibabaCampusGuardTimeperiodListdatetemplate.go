package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusGuardTimeperiodListdatetemplate 门禁控制器查询日期模版
// alibaba.campus.guard.timeperiod.listdatetemplate
//
// 门禁控制器查询日期模版
func AlibabaCampusGuardTimeperiodListdatetemplate(clt *core.SDKClient, req *campus.AlibabaCampusGuardTimeperiodListdatetemplateAPIRequest, session string) (*campus.AlibabaCampusGuardTimeperiodListdatetemplateAPIResponse, error) {
	var resp campus.AlibabaCampusGuardTimeperiodListdatetemplateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
