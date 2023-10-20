package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusGuardTimeperiodListdatetemplate 门禁控制器查询日期模版
// alibaba.campus.guard.timeperiod.listdatetemplate
//
// 门禁控制器查询日期模版
func AlibabaCampusGuardTimeperiodListdatetemplate(clt *core.SDKClient, req *campus.AlibabaCampusGuardTimeperiodListdatetemplateAPIRequest, resp *campus.AlibabaCampusGuardTimeperiodListdatetemplateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
