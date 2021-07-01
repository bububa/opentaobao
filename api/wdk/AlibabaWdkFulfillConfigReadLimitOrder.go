package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkFulfillConfigReadLimitOrder
根据仓code查询仓限单配置
alibaba.wdk.fulfill.config.read.limit.order

根据仓code查询仓限单配置 */
func AlibabaWdkFulfillConfigReadLimitOrder(clt *core.SDKClient, req *wdk.AlibabaWdkFulfillConfigReadLimitOrderAPIRequest, session string) (*wdk.AlibabaWdkFulfillConfigReadLimitOrderAPIResponse, error) {
	var resp wdk.AlibabaWdkFulfillConfigReadLimitOrderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
