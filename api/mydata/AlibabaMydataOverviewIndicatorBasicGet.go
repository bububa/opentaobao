package mydata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mydata"
)

// AlibabaMydataOverviewIndicatorBasicGet 我的效果-获取公司询盘流量行业表现
// alibaba.mydata.overview.indicator.basic.get
//
// 获取公司询盘流量行业表现
func AlibabaMydataOverviewIndicatorBasicGet(ctx context.Context, clt *core.SDKClient, req *mydata.AlibabaMydataOverviewIndicatorBasicGetAPIRequest, resp *mydata.AlibabaMydataOverviewIndicatorBasicGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
