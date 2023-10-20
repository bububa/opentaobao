package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse 根据交易单号判断是否为菜鸟发货订单 API返回值
// cainiao.cboss.workplatform.logistics.iscainiaoorder
//
// 根据交易单号判断是否为菜鸟发货订单
type CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse struct {
	model.CommonResponse
	CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponseModel).Reset()
}

// CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponseModel is 根据交易单号判断是否为菜鸟发货订单 成功返回结果
type CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cboss_workplatform_logistics_iscainiaoorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorCode
	ResErrorCode string `json:"res_error_code,omitempty" xml:"res_error_code,omitempty"`
	// errorMsg
	ResErrorMsg string `json:"res_error_msg,omitempty" xml:"res_error_msg,omitempty"`
	// isCainiaoOrder
	IsCainiaoOrder bool `json:"is_cainiao_order,omitempty" xml:"is_cainiao_order,omitempty"`
	// success
	ResSuccess bool `json:"res_success,omitempty" xml:"res_success,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResErrorCode = ""
	m.ResErrorMsg = ""
	m.IsCainiaoOrder = false
	m.ResSuccess = false
}

var poolCainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse)
	},
}

// GetCainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse 从 sync.Pool 获取 CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse
func GetCainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse() *CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse {
	return poolCainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse.Get().(*CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse)
}

// ReleaseCainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse 将 CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse 保存到 sync.Pool
func ReleaseCainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse(v *CainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse) {
	v.Reset()
	poolCainiaoCbossWorkplatformLogisticsIscainiaoorderAPIResponse.Put(v)
}
