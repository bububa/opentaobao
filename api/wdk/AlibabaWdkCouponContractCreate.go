package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkcouponcontractcreate 营销券合同创建接口
// alibaba.wdk.coupon.contract.create
//
// 营销券合同创建接口
func Alibabawdkcouponcontractcreate(clt *core.SDKClient, req *wdk.AlibabawdkcouponcontractcreateAPIRequest, session string) (*wdk.AlibabawdkcouponcontractcreateAPIResponse, error) {
	var resp wdk.AlibabawdkcouponcontractcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
