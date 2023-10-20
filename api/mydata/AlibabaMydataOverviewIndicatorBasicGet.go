package mydata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mydata"
)

// Alibabamydataoverviewindicatorbasicget 我的效果-获取公司询盘流量行业表现
// alibaba.mydata.overview.indicator.basic.get
//
// 获取公司询盘流量行业表现
func Alibabamydataoverviewindicatorbasicget(clt *core.SDKClient, req *mydata.AlibabamydataoverviewindicatorbasicgetAPIRequest, session string) (*mydata.AlibabamydataoverviewindicatorbasicgetAPIResponse, error) {
	var resp mydata.AlibabamydataoverviewindicatorbasicgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
