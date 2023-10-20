package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuUpdate 更新商品
// alibaba.wdk.sku.update
//
// 开放商品更新接口
func AlibabaWdkSkuUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkSkuUpdateAPIRequest, resp *wdk.AlibabaWdkSkuUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
