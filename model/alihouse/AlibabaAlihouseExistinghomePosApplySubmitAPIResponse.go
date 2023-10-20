package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomePosApplySubmitAPIResponse pos申请接口 API返回值
// alibaba.alihouse.existinghome.pos.apply.submit
//
// pos申请接口
type AlibabaAlihouseExistinghomePosApplySubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomePosApplySubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomePosApplySubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomePosApplySubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomePosApplySubmitAPIResponseModel is pos申请接口 成功返回结果
type AlibabaAlihouseExistinghomePosApplySubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_pos_apply_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 对象
	Result *AlibabaAlihouseExistinghomePosApplySubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomePosApplySubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomePosApplySubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomePosApplySubmitAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomePosApplySubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomePosApplySubmitAPIResponse
func GetAlibabaAlihouseExistinghomePosApplySubmitAPIResponse() *AlibabaAlihouseExistinghomePosApplySubmitAPIResponse {
	return poolAlibabaAlihouseExistinghomePosApplySubmitAPIResponse.Get().(*AlibabaAlihouseExistinghomePosApplySubmitAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomePosApplySubmitAPIResponse 将 AlibabaAlihouseExistinghomePosApplySubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomePosApplySubmitAPIResponse(v *AlibabaAlihouseExistinghomePosApplySubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomePosApplySubmitAPIResponse.Put(v)
}
