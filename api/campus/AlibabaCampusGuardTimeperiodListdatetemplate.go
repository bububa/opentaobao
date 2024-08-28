package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusGuardTimeperiodListdatetemplate 门禁控制器查询日期模版
// alibaba.campus.guard.timeperiod.listdatetemplate
//
// 门禁控制器查询日期模版
func AlibabaCampusGuardTimeperiodListdatetemplate(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusGuardTimeperiodListdatetemplateAPIRequest, resp *campus.AlibabaCampusGuardTimeperiodListdatetemplateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
