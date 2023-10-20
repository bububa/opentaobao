package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmMenuListAPIResponse 获取特价菜单 API返回值
// alibaba.alsc.crm.menu.list
//
// 获取特价菜单
type AlibabaAlscCrmMenuListAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmMenuListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmMenuListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmMenuListAPIResponseModel).Reset()
}

// AlibabaAlscCrmMenuListAPIResponseModel is 获取特价菜单 成功返回结果
type AlibabaAlscCrmMenuListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_menu_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页返回模型
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmMenuListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmMenuListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmMenuListAPIResponse)
	},
}

// GetAlibabaAlscCrmMenuListAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmMenuListAPIResponse
func GetAlibabaAlscCrmMenuListAPIResponse() *AlibabaAlscCrmMenuListAPIResponse {
	return poolAlibabaAlscCrmMenuListAPIResponse.Get().(*AlibabaAlscCrmMenuListAPIResponse)
}

// ReleaseAlibabaAlscCrmMenuListAPIResponse 将 AlibabaAlscCrmMenuListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmMenuListAPIResponse(v *AlibabaAlscCrmMenuListAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmMenuListAPIResponse.Put(v)
}
