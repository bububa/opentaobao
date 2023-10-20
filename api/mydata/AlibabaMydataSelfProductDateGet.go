package mydata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mydata"
)

// Alibabamydataselfproductdateget 获取客户产品相关表现数据的可用时间范围
// alibaba.mydata.self.product.date.get
//
// 获取客户产品相关表现数据的可用时间范围
func Alibabamydataselfproductdateget(clt *core.SDKClient, req *mydata.AlibabamydataselfproductdategetAPIRequest, session string) (*mydata.AlibabamydataselfproductdategetAPIResponse, error) {
	var resp mydata.AlibabamydataselfproductdategetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
