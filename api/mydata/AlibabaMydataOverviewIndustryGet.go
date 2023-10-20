package mydata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mydata"
)

// Alibabamydataoverviewindustryget 我的效果-获取Top行业列表
// alibaba.mydata.overview.industry.get
//
// 获取数据管家我的效果API可以使用的行业
func Alibabamydataoverviewindustryget(clt *core.SDKClient, req *mydata.AlibabamydataoverviewindustrygetAPIRequest, session string) (*mydata.AlibabamydataoverviewindustrygetAPIResponse, error) {
	var resp mydata.AlibabamydataoverviewindustrygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
