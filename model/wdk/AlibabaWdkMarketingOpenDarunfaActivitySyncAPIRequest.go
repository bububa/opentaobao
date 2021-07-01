package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingOpenDarunfaActivitySyncAPIRequest
活动数据同步 API请求
alibaba.wdk.marketing.open.darunfa.activity.sync

大润发活动数据同步 */
type AlibabaWdkMarketingOpenDarunfaActivitySyncAPIRequest struct {
	model.Params
	// 大润发活动数据
	_activityList []DrfTxdActivityBo
	// 门店Id
	_shopId string
	// 版本ID
	_versionId int64
}

// New
