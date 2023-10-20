package aeusergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aeusergrowth"
)

// Aliexpressusergrowthsearchitemsget 第三方平台搜索AE商品
// aliexpress.usergrowth.search.items.get
//
// 第三方平台的搜索服务   获取AE商品list
func Aliexpressusergrowthsearchitemsget(clt *core.SDKClient, req *aeusergrowth.AliexpressusergrowthsearchitemsgetAPIRequest, session string) (*aeusergrowth.AliexpressusergrowthsearchitemsgetAPIResponse, error) {
	var resp aeusergrowth.AliexpressusergrowthsearchitemsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
