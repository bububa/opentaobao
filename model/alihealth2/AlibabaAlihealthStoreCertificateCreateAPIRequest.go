package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthStoreCertificateCreateAPIRequest
仓库换证审批 API请求
alibaba.alihealth.store.certificate.create

仓库侧换证发起审批 */
type AlibabaAlihealthStoreCertificateCreateAPIRequest struct {
	model.Params
	// 仓库code
	_storeCode string
	// 审批业务类型
	_auditType string
	// 审批内容
	_content string
	// 业务单号
	_bizNo string
}

// New
