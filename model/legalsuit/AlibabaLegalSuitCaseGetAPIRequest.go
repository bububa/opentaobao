package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalSuitCaseGetAPIRequest
获取案件信息接口v2版本 API请求
alibaba.legal.suit.case.get

获取案件信息 */
type AlibabaLegalSuitCaseGetAPIRequest struct {
	model.Params
	// 案件id
	_id int64
}

// New
