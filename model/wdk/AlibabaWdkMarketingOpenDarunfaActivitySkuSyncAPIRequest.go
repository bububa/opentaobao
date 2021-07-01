package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest
营销商品数据同步 API请求
alibaba.wdk.marketing.open.darunfa.activity.sku.sync

大润发营销商品数据同步 */
type AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIRequest struct {
	model.Params
	// 淘鲜达活动商品信息
	_activitySkuList []DrfTxdActivitySkuBo
	// 数据版本Id
	_versionId int64
	// 大润发活动Id
	_activityId string
	// 活动对应的门店Id
	_shopId string
}

// New
