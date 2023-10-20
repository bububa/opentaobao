package aecreatives

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// AliexpressAffiliateFeaturedpromoGet 联盟主题推广活动信息获取
// aliexpress.affiliate.featuredpromo.get
//
// 获取联盟主题推广活动信息
func AliexpressAffiliateFeaturedpromoGet(clt *core.SDKClient, req *aecreatives.AliexpressAffiliateFeaturedpromoGetAPIRequest, resp *aecreatives.AliexpressAffiliateFeaturedpromoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
