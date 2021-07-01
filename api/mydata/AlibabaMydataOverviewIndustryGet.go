package mydata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mydata"
)

/* AlibabaMydataOverviewIndustryGet
我的效果-获取Top行业列表
alibaba.mydata.overview.industry.get

获取数据管家我的效果API可以使用的行业 */
func AlibabaMydataOverviewIndustryGet(clt *core.SDKClient, req *mydata.AlibabaMydataOverviewIndustryGetAPIRequest, session string) (*mydata.AlibabaMydataOverviewIndustryGetAPIResponse, error) {
	var resp mydata.AlibabaMydataOverviewIndustryGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
