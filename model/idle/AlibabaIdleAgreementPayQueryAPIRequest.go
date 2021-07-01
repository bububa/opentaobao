package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleAgreementPayQueryAPIRequest
代扣详情查询 API请求
alibaba.idle.agreement.pay.query

查询代扣结果 */
type AlibabaIdleAgreementPayQueryAPIRequest struct {
	model.Params
	// 入参
	_param *AgreementPayBillQueryParam
}

// New
