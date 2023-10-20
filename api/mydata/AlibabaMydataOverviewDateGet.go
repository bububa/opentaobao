package mydata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mydata"
)

// Alibabamydataoverviewdateget 我的效果-获取数据周期
// alibaba.mydata.overview.date.get
//
// 获取数据管家我的效果API可以使用的数据周期
func Alibabamydataoverviewdateget(clt *core.SDKClient, req *mydata.AlibabamydataoverviewdategetAPIRequest, session string) (*mydata.AlibabamydataoverviewdategetAPIResponse, error) {
	var resp mydata.AlibabamydataoverviewdategetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
