package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingVersionGenerateAPIRequest
生成发布使用的版本号 API请求
alibaba.wdk.marketing.version.generate

生成发布使用的版本号 */
type AlibabaWdkMarketingVersionGenerateAPIRequest struct {
	model.Params
	// 档期版本号参数信息
	_param *SeasonVersionParam
}

// New
