package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingVersionCommitAPIRequest
提交版本号 API请求
alibaba.wdk.marketing.version.commit

提交版本号，标识结束此版本操作 */
type AlibabaWdkMarketingVersionCommitAPIRequest struct {
	model.Params
	// 版本号提交参数
	_param *SeasonVersionCommitParam
}

// New
