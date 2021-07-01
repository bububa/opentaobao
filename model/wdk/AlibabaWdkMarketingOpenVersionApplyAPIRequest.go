package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingOpenVersionApplyAPIRequest
数据同步版本号申请 API请求
alibaba.wdk.marketing.open.version.apply

数据同步版本号申请 */
type AlibabaWdkMarketingOpenVersionApplyAPIRequest struct {
	model.Params
	// 同步版本信息
	_syncVersion *SyncVersionBo
}

// New
