package aecreatives

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// Aliexpressaffiliatefeaturedpromoget 联盟主题推广活动信息获取
// aliexpress.affiliate.featuredpromo.get
//
// 获取联盟主题推广活动信息
func Aliexpressaffiliatefeaturedpromoget(clt *core.SDKClient, req *aecreatives.AliexpressaffiliatefeaturedpromogetAPIRequest, session string) (*aecreatives.AliexpressaffiliatefeaturedpromogetAPIResponse, error) {
	var resp aecreatives.AliexpressaffiliatefeaturedpromogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
