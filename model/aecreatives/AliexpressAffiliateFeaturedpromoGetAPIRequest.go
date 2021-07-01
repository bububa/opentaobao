package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAffiliateFeaturedpromoGetAPIRequest
联盟主题推广活动信息获取 API请求
aliexpress.affiliate.featuredpromo.get

获取联盟主题推广活动信息 */
type AliexpressAffiliateFeaturedpromoGetAPIRequest struct {
	model.Params
	// 请求签名
	_appSignature string
	// 返回字段列表
	_fields string
}

// New
