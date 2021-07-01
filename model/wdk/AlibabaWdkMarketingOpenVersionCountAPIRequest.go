package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingOpenVersionCountAPIRequest
版本数量查询 API请求
alibaba.wdk.marketing.open.version.count

版本数量查询 */
type AlibabaWdkMarketingOpenVersionCountAPIRequest struct {
	model.Params
	// 查询版本号
	_versionId int64
	// 操作Id
	_operateId string
}

// New
