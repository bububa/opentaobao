package wms

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCrmOmsRuleSyncAPIResponse 商家ERP订单处理规则同步 API返回值
// cainiao.crm.oms.rule.sync
//
// 将商家ERP订单处理规则同步到菜鸟CRM系统
type CainiaoCrmOmsRuleSyncAPIResponse struct {
	model.CommonResponse
	CainiaoCrmOmsRuleSyncAPIResponseModel
}

// CainiaoCrmOmsRuleSyncAPIResponseModel is 商家ERP订单处理规则同步 成功返回结果
type CainiaoCrmOmsRuleSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_crm_oms_rule_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorCode
	WlErrorCode string `json:"wl_error_code,omitempty" xml:"wl_error_code,omitempty"`
	// errorMsg
	WlErrorMsg string `json:"wl_error_msg,omitempty" xml:"wl_error_msg,omitempty"`
	// success
	WlSuccess bool `json:"wl_success,omitempty" xml:"wl_success,omitempty"`
}
