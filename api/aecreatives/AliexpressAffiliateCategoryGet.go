package aecreatives

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// Aliexpressaffiliatecategoryget AE流量推广类目信息获取API
// aliexpress.affiliate.category.get
//
// 获取AE流量推广类目的API
func Aliexpressaffiliatecategoryget(clt *core.SDKClient, req *aecreatives.AliexpressaffiliatecategorygetAPIRequest, session string) (*aecreatives.AliexpressaffiliatecategorygetAPIResponse, error) {
	var resp aecreatives.AliexpressaffiliatecategorygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
