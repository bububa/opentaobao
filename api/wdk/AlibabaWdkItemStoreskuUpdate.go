package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkItemStoreskuUpdate 五道口商品中心门店商品修改
// alibaba.wdk.item.storesku.update
//
// 五道口商品中心门店商品修改
func AlibabaWdkItemStoreskuUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkItemStoreskuUpdateAPIRequest, resp *wdk.AlibabaWdkItemStoreskuUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
