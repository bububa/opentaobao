package mydata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mydata"
)

// AlibabaMydataSelfProductGet 获取客户产品相关表现数据
// alibaba.mydata.self.product.get
//
// 获取客户产品相关表现数据
func AlibabaMydataSelfProductGet(ctx context.Context, clt *core.SDKClient, req *mydata.AlibabaMydataSelfProductGetAPIRequest, resp *mydata.AlibabaMydataSelfProductGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
