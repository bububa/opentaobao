package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdlePayPlanCreateAPIRequest
创建代扣计划 API请求
alibaba.idle.pay.plan.create

闲鱼平台代扣能力：
1、用户和闲鱼签约代扣协议 服务商通过直付通产品挂载成为闲鱼二级商户 来完成用户和服务商的结算
2、创建代扣计划 */
type AlibabaIdlePayPlanCreateAPIRequest struct {
	model.Params
	// 业务入参
	_agreementPayPlanParam *AgreementPayPlanParam
}

// New
