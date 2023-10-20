package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse 菜鸟工单平台根据业务类型id查询业务类型详细信息 API返回值
// cainiao.cboss.workplatform.biztype.querybyid
//
// 菜鸟工单平台根据业务类型id查询业务类型详细信息。 目前调用者ISV
type CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse struct {
	model.CommonResponse
	CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponseModel).Reset()
}

// CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponseModel is 菜鸟工单平台根据业务类型id查询业务类型详细信息 成功返回结果
type CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cboss_workplatform_biztype_querybyid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CainiaoCbossWorkplatformBiztypeQuerybyidResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse)
	},
}

// GetCainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse 从 sync.Pool 获取 CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse
func GetCainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse() *CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse {
	return poolCainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse.Get().(*CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse)
}

// ReleaseCainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse 将 CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse 保存到 sync.Pool
func ReleaseCainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse(v *CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse) {
	v.Reset()
	poolCainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse.Put(v)
}
