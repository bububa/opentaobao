package aeusergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aeusergrowth"
)

// AliexpressUsergrowthSearchItemsGet 第三方平台搜索AE商品
// aliexpress.usergrowth.search.items.get
//
// 第三方平台的搜索服务   获取AE商品list
func AliexpressUsergrowthSearchItemsGet(clt *core.SDKClient, req *aeusergrowth.AliexpressUsergrowthSearchItemsGetAPIRequest, session string) (*aeusergrowth.AliexpressUsergrowthSearchItemsGetAPIResponse, error) {
	var resp aeusergrowth.AliexpressUsergrowthSearchItemsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
