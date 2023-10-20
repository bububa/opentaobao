package aedropshiper

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressDsImageSearchAPIResponse 图片搜索 API返回值
// aliexpress.ds.image.search
//
// 图片搜索
type AliexpressDsImageSearchAPIResponse struct {
	model.CommonResponse
	AliexpressDsImageSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressDsImageSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressDsImageSearchAPIResponseModel).Reset()
}

// AliexpressDsImageSearchAPIResponseModel is 图片搜索 成功返回结果
type AliexpressDsImageSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ds_image_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// error code
	RspCode string `json:"rsp_code,omitempty" xml:"rsp_code,omitempty"`
	// System Error
	RspMsg string `json:"rsp_msg,omitempty" xml:"rsp_msg,omitempty"`
	// result object
	Data *TrafficImageSearchResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// total record count
	TotalRecordCount int64 `json:"total_record_count,omitempty" xml:"total_record_count,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressDsImageSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RspCode = ""
	m.RspMsg = ""
	m.Data = nil
	m.TotalRecordCount = 0
}

var poolAliexpressDsImageSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressDsImageSearchAPIResponse)
	},
}

// GetAliexpressDsImageSearchAPIResponse 从 sync.Pool 获取 AliexpressDsImageSearchAPIResponse
func GetAliexpressDsImageSearchAPIResponse() *AliexpressDsImageSearchAPIResponse {
	return poolAliexpressDsImageSearchAPIResponse.Get().(*AliexpressDsImageSearchAPIResponse)
}

// ReleaseAliexpressDsImageSearchAPIResponse 将 AliexpressDsImageSearchAPIResponse 保存到 sync.Pool
func ReleaseAliexpressDsImageSearchAPIResponse(v *AliexpressDsImageSearchAPIResponse) {
	v.Reset()
	poolAliexpressDsImageSearchAPIResponse.Put(v)
}
