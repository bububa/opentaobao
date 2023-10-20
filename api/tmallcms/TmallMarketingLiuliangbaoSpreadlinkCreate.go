package tmallcms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcms"
)

// TmallMarketingLiuliangbaoSpreadlinkCreate 创建流量宝活动链接
// tmall.marketing.liuliangbao.spreadlink.create
//
// 通过源活动链接和商家ID，创建流量宝活动链接
func TmallMarketingLiuliangbaoSpreadlinkCreate(clt *core.SDKClient, req *tmallcms.TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest, resp *tmallcms.TmallMarketingLiuliangbaoSpreadlinkCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
