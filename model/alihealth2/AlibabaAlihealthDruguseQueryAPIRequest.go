package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDruguseQueryAPIRequest
合理用药规则查询 API请求
alibaba.alihealth.druguse.query

查询用户购买的药品命中的风险规则 */
type AlibabaAlihealthDruguseQueryAPIRequest struct {
	model.Params
	// 入参
	_command *SafeMedicationReqCommand
}

// New
