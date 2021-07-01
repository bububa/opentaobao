package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlifanyiMarketLoginAPIRequest
登陆用户 API请求
alibaba.alifanyi.market.login

企业或组织购买软件服务后可登陆阿里翻译众包系统，接口返回该企业的用户。 */
type AlibabaAlifanyiMarketLoginAPIRequest struct {
	model.Params
	// 请求参数
	_reportQueryApiDTO *ReportQueryApiDto
}

// New
