package mydata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mydata"
)

// AlibabaMydataSelfProductDateGet 获取客户产品相关表现数据的可用时间范围
// alibaba.mydata.self.product.date.get
//
// 获取客户产品相关表现数据的可用时间范围
func AlibabaMydataSelfProductDateGet(clt *core.SDKClient, req *mydata.AlibabaMydataSelfProductDateGetAPIRequest, resp *mydata.AlibabaMydataSelfProductDateGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
