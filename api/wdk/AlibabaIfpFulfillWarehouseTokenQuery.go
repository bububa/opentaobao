package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaIfpFulfillWarehouseTokenQuery
同城令牌打印接口
alibaba.ifp.fulfill.warehouse.token.query

用于仓内作业打印包裹信息 */
func AlibabaIfpFulfillWarehouseTokenQuery(clt *core.SDKClient, req *wdk.AlibabaIfpFulfillWarehouseTokenQueryAPIRequest, session string) (*wdk.AlibabaIfpFulfillWarehouseTokenQueryAPIResponse, error) {
	var resp wdk.AlibabaIfpFulfillWarehouseTokenQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
