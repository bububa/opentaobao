package baoxian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlipayBaoxianClaimUpdateAPIRequest
更新赔案 API请求
alipay.baoxian.claim.update

更新保险理赔单 */
type AlipayBaoxianClaimUpdateAPIRequest struct {
	model.Params
	// 业务数据
	_bizData string
	// 进度列表
	_progressList []string
	// 附件列表
	_claimAttachments []ClaimAttachment
	// 保单业务单号
	_policyBizNo string
	// 外部业务单号
	_outBizNo string
	// 业务来源
	_bizSource string
	// 理赔金额(单位为分)
	_claimFee int64
	// 理赔单号
	_claimNo string
	// 理赔外部业务单号
	_claimOutBizNo string
	// 标准产品ID
	_spNo string
}

// New
