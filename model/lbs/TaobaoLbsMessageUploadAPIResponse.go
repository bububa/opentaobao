package lbs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLbsMessageUploadAPIResponse lbs数据采集 API返回值
// taobao.lbs.message.upload
//
// lbs数据采集
type TaobaoLbsMessageUploadAPIResponse struct {
	model.CommonResponse
	TaobaoLbsMessageUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLbsMessageUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLbsMessageUploadAPIResponseModel).Reset()
}

// TaobaoLbsMessageUploadAPIResponseModel is lbs数据采集 成功返回结果
type TaobaoLbsMessageUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"lbs_message_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// subCode
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// subMsg
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// result
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLbsMessageUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = ""
	m.Result = false
}

var poolTaobaoLbsMessageUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLbsMessageUploadAPIResponse)
	},
}

// GetTaobaoLbsMessageUploadAPIResponse 从 sync.Pool 获取 TaobaoLbsMessageUploadAPIResponse
func GetTaobaoLbsMessageUploadAPIResponse() *TaobaoLbsMessageUploadAPIResponse {
	return poolTaobaoLbsMessageUploadAPIResponse.Get().(*TaobaoLbsMessageUploadAPIResponse)
}

// ReleaseTaobaoLbsMessageUploadAPIResponse 将 TaobaoLbsMessageUploadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLbsMessageUploadAPIResponse(v *TaobaoLbsMessageUploadAPIResponse) {
	v.Reset()
	poolTaobaoLbsMessageUploadAPIResponse.Put(v)
}
