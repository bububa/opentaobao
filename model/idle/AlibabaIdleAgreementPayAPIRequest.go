package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleAgreementPayAPIRequest
闲鱼平台商户代扣 API请求
alibaba.idle.agreement.pay

闲鱼平台代扣能力：用户和闲鱼签约代扣协议 服务商通过直付通产品挂载成为闲鱼二级商户 来完成用户和服务商的结算 */
type AlibabaIdleAgreementPayAPIRequest struct {
	model.Params
	// 协议代扣参数
	_agreementPayParam *AgreementPayParam
}

// New
