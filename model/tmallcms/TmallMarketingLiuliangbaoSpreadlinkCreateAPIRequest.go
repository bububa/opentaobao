package tmallcms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest
创建流量宝活动链接 API请求
tmall.marketing.liuliangbao.spreadlink.create

通过源活动链接和商家ID，创建流量宝活动链接 */
type TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest struct {
	model.Params
	// 活动链接，必须为淘系链接
	_url string
}

// New
