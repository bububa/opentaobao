package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthPwApplynodeUpdatenameAPIRequest
回调变更患者姓名 API请求
alibaba.alihealth.pw.applynode.updatename

回调变更患者姓名 */
type AlibabaAlihealthPwApplynodeUpdatenameAPIRequest struct {
	model.Params
	// 回调入参
	_body *ModifyNameRo
}

// New
