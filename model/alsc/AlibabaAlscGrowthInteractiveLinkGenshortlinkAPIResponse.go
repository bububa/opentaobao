package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse 短链接口 API返回值
// alibaba.alsc.growth.interactive.link.genshortlink
//
// 短链接口
type AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse struct {
	model.CommonResponse
	AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponseModel).Reset()
}

// AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponseModel is 短链接口 成功返回结果
type AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_growth_interactive_link_genshortlink_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *BaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse)
	},
}

// GetAlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse 从 sync.Pool 获取 AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse
func GetAlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse() *AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse {
	return poolAlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse.Get().(*AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse)
}

// ReleaseAlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse 将 AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse(v *AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse) {
	v.Reset()
	poolAlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse.Put(v)
}
