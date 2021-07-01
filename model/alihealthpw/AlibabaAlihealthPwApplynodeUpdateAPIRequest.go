package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthPwApplynodeUpdateAPIRequest
申请节点变更回调 API请求
alibaba.alihealth.pw.applynode.update

基金会回调阿里健康更新药品援助申请单的状态 */
type AlibabaAlihealthPwApplynodeUpdateAPIRequest struct {
	model.Params
	// 回调入参
	_body *AuditRollbackRo
}

// New
