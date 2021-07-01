package mydata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mydata"
)

/* AlibabaMydataSelfProductGet
获取客户产品相关表现数据
alibaba.mydata.self.product.get

获取客户产品相关表现数据 */
func AlibabaMydataSelfProductGet(clt *core.SDKClient, req *mydata.AlibabaMydataSelfProductGetAPIRequest, session string) (*mydata.AlibabaMydataSelfProductGetAPIResponse, error) {
	var resp mydata.AlibabaMydataSelfProductGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
