package mydata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mydata"
)

// AlibabaMydataOverviewDateGet 我的效果-获取数据周期
// alibaba.mydata.overview.date.get
//
// 获取数据管家我的效果API可以使用的数据周期
func AlibabaMydataOverviewDateGet(ctx context.Context, clt *core.SDKClient, req *mydata.AlibabaMydataOverviewDateGetAPIRequest, resp *mydata.AlibabaMydataOverviewDateGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
