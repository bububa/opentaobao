package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmActivityGetbaseinfoAPIResponse ISV查询活动 API返回值
// alibaba.lsy.crm.activity.getbaseinfo
//
// ISV查询活动
type AlibabaLsyCrmActivityGetbaseinfoAPIResponse struct {
	model.CommonResponse
	AlibabaLsyCrmActivityGetbaseinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLsyCrmActivityGetbaseinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLsyCrmActivityGetbaseinfoAPIResponseModel).Reset()
}

// AlibabaLsyCrmActivityGetbaseinfoAPIResponseModel is ISV查询活动 成功返回结果
type AlibabaLsyCrmActivityGetbaseinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_activity_getbaseinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaLsyCrmActivityGetbaseinfoResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLsyCrmActivityGetbaseinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLsyCrmActivityGetbaseinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLsyCrmActivityGetbaseinfoAPIResponse)
	},
}

// GetAlibabaLsyCrmActivityGetbaseinfoAPIResponse 从 sync.Pool 获取 AlibabaLsyCrmActivityGetbaseinfoAPIResponse
func GetAlibabaLsyCrmActivityGetbaseinfoAPIResponse() *AlibabaLsyCrmActivityGetbaseinfoAPIResponse {
	return poolAlibabaLsyCrmActivityGetbaseinfoAPIResponse.Get().(*AlibabaLsyCrmActivityGetbaseinfoAPIResponse)
}

// ReleaseAlibabaLsyCrmActivityGetbaseinfoAPIResponse 将 AlibabaLsyCrmActivityGetbaseinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLsyCrmActivityGetbaseinfoAPIResponse(v *AlibabaLsyCrmActivityGetbaseinfoAPIResponse) {
	v.Reset()
	poolAlibabaLsyCrmActivityGetbaseinfoAPIResponse.Put(v)
}
