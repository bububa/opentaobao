package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSeriesSkuRemove 系列品商品变更-移除商品
// alibaba.wdk.series.sku.remove
//
// 系列品商品变更-移除商品
func AlibabaWdkSeriesSkuRemove(clt *core.SDKClient, req *wdk.AlibabaWdkSeriesSkuRemoveAPIRequest, resp *wdk.AlibabaWdkSeriesSkuRemoveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
