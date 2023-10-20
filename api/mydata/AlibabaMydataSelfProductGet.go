package mydata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mydata"
)

// Alibabamydataselfproductget 获取客户产品相关表现数据
// alibaba.mydata.self.product.get
//
// 获取客户产品相关表现数据
func Alibabamydataselfproductget(clt *core.SDKClient, req *mydata.AlibabamydataselfproductgetAPIRequest, session string) (*mydata.AlibabamydataselfproductgetAPIResponse, error) {
	var resp mydata.AlibabamydataselfproductgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
