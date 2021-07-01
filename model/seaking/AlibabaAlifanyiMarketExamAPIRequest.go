package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlifanyiMarketExamAPIRequest
通过考试用户 API请求
alibaba.alifanyi.market.exam

企业或组织购买软件服务后可参与阿里翻译在线系统的考试认证，接口返回该企业或组织认证通过的用户 */
type AlibabaAlifanyiMarketExamAPIRequest struct {
	model.Params
	// 请求参数
	_reportQueryApiDTO *ReportQueryApiDto
}

// New
