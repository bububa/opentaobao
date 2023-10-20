package tmalltrend

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallBrandItemUploadAPIResponse 天猫品牌新品同步API API返回值
// tmall.brand.item.upload
//
// 支撑天猫品牌将各渠道新品信息同步至平台
type TmallBrandItemUploadAPIResponse struct {
	model.CommonResponse
	TmallBrandItemUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TmallBrandItemUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallBrandItemUploadAPIResponseModel).Reset()
}

// TmallBrandItemUploadAPIResponseModel is 天猫品牌新品同步API 成功返回结果
type TmallBrandItemUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_brand_item_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 请求参数错误
	RespErrorCode int64 `json:"resp_error_code,omitempty" xml:"resp_error_code,omitempty"`
	// 是否成功
	RespSuccess bool `json:"resp_success,omitempty" xml:"resp_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallBrandItemUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Value = ""
	m.ErrorMsg = ""
	m.RespErrorCode = 0
	m.RespSuccess = false
}

var poolTmallBrandItemUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallBrandItemUploadAPIResponse)
	},
}

// GetTmallBrandItemUploadAPIResponse 从 sync.Pool 获取 TmallBrandItemUploadAPIResponse
func GetTmallBrandItemUploadAPIResponse() *TmallBrandItemUploadAPIResponse {
	return poolTmallBrandItemUploadAPIResponse.Get().(*TmallBrandItemUploadAPIResponse)
}

// ReleaseTmallBrandItemUploadAPIResponse 将 TmallBrandItemUploadAPIResponse 保存到 sync.Pool
func ReleaseTmallBrandItemUploadAPIResponse(v *TmallBrandItemUploadAPIResponse) {
	v.Reset()
	poolTmallBrandItemUploadAPIResponse.Put(v)
}
