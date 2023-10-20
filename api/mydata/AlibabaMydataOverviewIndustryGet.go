package mydata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mydata"
)

// AlibabaMydataOverviewIndustryGet 我的效果-获取Top行业列表
// alibaba.mydata.overview.industry.get
//
// 获取数据管家我的效果API可以使用的行业
func AlibabaMydataOverviewIndustryGet(clt *core.SDKClient, req *mydata.AlibabaMydataOverviewIndustryGetAPIRequest, resp *mydata.AlibabaMydataOverviewIndustryGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
