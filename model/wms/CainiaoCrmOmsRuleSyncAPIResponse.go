package wms

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *CainiaoCrmOmsRuleSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoCrmOmsRuleSyncAPIResponseModel).Reset()
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

// Reset 清空结构体
func (m *CainiaoCrmOmsRuleSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WlErrorCode = ""
	m.WlErrorMsg = ""
	m.WlSuccess = false
}

var poolCainiaoCrmOmsRuleSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoCrmOmsRuleSyncAPIResponse)
	},
}

// GetCainiaoCrmOmsRuleSyncAPIResponse 从 sync.Pool 获取 CainiaoCrmOmsRuleSyncAPIResponse
func GetCainiaoCrmOmsRuleSyncAPIResponse() *CainiaoCrmOmsRuleSyncAPIResponse {
	return poolCainiaoCrmOmsRuleSyncAPIResponse.Get().(*CainiaoCrmOmsRuleSyncAPIResponse)
}

// ReleaseCainiaoCrmOmsRuleSyncAPIResponse 将 CainiaoCrmOmsRuleSyncAPIResponse 保存到 sync.Pool
func ReleaseCainiaoCrmOmsRuleSyncAPIResponse(v *CainiaoCrmOmsRuleSyncAPIResponse) {
	v.Reset()
	poolCainiaoCrmOmsRuleSyncAPIResponse.Put(v)
}
