package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsMediaResourcesUploadAPIResponse 图片与视频上传 API返回值
// taobao.logistics.media.resources.upload
//
// 图片与视频上传
type TaobaoLogisticsMediaResourcesUploadAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsMediaResourcesUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsMediaResourcesUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsMediaResourcesUploadAPIResponseModel).Reset()
}

// TaobaoLogisticsMediaResourcesUploadAPIResponseModel is 图片与视频上传 成功返回结果
type TaobaoLogisticsMediaResourcesUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_media_resources_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 图片视频上传出参
	Result *LspTopResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsMediaResourcesUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoLogisticsMediaResourcesUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsMediaResourcesUploadAPIResponse)
	},
}

// GetTaobaoLogisticsMediaResourcesUploadAPIResponse 从 sync.Pool 获取 TaobaoLogisticsMediaResourcesUploadAPIResponse
func GetTaobaoLogisticsMediaResourcesUploadAPIResponse() *TaobaoLogisticsMediaResourcesUploadAPIResponse {
	return poolTaobaoLogisticsMediaResourcesUploadAPIResponse.Get().(*TaobaoLogisticsMediaResourcesUploadAPIResponse)
}

// ReleaseTaobaoLogisticsMediaResourcesUploadAPIResponse 将 TaobaoLogisticsMediaResourcesUploadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsMediaResourcesUploadAPIResponse(v *TaobaoLogisticsMediaResourcesUploadAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsMediaResourcesUploadAPIResponse.Put(v)
}
