package aeusergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aeusergrowth"
)

// AliexpressUsergrowthSearchItemsGet 第三方平台搜索AE商品
// aliexpress.usergrowth.search.items.get
//
// 第三方平台的搜索服务   获取AE商品list
func AliexpressUsergrowthSearchItemsGet(clt *core.SDKClient, req *aeusergrowth.AliexpressUsergrowthSearchItemsGetAPIRequest, resp *aeusergrowth.AliexpressUsergrowthSearchItemsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
