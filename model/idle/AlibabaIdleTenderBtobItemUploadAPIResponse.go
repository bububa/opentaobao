package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTenderBtobItemUploadAPIResponse 暗拍发布/编辑B2B商品 API返回值
// alibaba.idle.tender.btob.item.upload
//
// 暗拍发布/编辑B2B商品
type AlibabaIdleTenderBtobItemUploadAPIResponse struct {
	model.CommonResponse
	AlibabaIdleTenderBtobItemUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleTenderBtobItemUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleTenderBtobItemUploadAPIResponseModel).Reset()
}

// AlibabaIdleTenderBtobItemUploadAPIResponseModel is 暗拍发布/编辑B2B商品 成功返回结果
type AlibabaIdleTenderBtobItemUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_tender_btob_item_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleTenderBtobItemUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleTenderBtobItemUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleTenderBtobItemUploadAPIResponse)
	},
}

// GetAlibabaIdleTenderBtobItemUploadAPIResponse 从 sync.Pool 获取 AlibabaIdleTenderBtobItemUploadAPIResponse
func GetAlibabaIdleTenderBtobItemUploadAPIResponse() *AlibabaIdleTenderBtobItemUploadAPIResponse {
	return poolAlibabaIdleTenderBtobItemUploadAPIResponse.Get().(*AlibabaIdleTenderBtobItemUploadAPIResponse)
}

// ReleaseAlibabaIdleTenderBtobItemUploadAPIResponse 将 AlibabaIdleTenderBtobItemUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleTenderBtobItemUploadAPIResponse(v *AlibabaIdleTenderBtobItemUploadAPIResponse) {
	v.Reset()
	poolAlibabaIdleTenderBtobItemUploadAPIResponse.Put(v)
}
