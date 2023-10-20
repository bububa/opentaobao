package tmalltrend

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTrendStyleProduceinfoUploadAPIResponse 款式生产信息同步API API返回值
// tmall.trend.style.produceinfo.upload
//
// 款式生产信息同步至平台
type TmallTrendStyleProduceinfoUploadAPIResponse struct {
	model.CommonResponse
	TmallTrendStyleProduceinfoUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTrendStyleProduceinfoUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTrendStyleProduceinfoUploadAPIResponseModel).Reset()
}

// TmallTrendStyleProduceinfoUploadAPIResponseModel is 款式生产信息同步API 成功返回结果
type TmallTrendStyleProduceinfoUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_trend_style_produceinfo_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 同步款式生产信息处理结果描述
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 错误描述信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码,1001-系统错误，1002-请求参数错误，1003-业务处理错误，1004-数据权限错误
	RespErrorCode int64 `json:"resp_error_code,omitempty" xml:"resp_error_code,omitempty"`
	// 是否成功
	RespSuccess bool `json:"resp_success,omitempty" xml:"resp_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallTrendStyleProduceinfoUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Value = ""
	m.ErrorMsg = ""
	m.RespErrorCode = 0
	m.RespSuccess = false
}

var poolTmallTrendStyleProduceinfoUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTrendStyleProduceinfoUploadAPIResponse)
	},
}

// GetTmallTrendStyleProduceinfoUploadAPIResponse 从 sync.Pool 获取 TmallTrendStyleProduceinfoUploadAPIResponse
func GetTmallTrendStyleProduceinfoUploadAPIResponse() *TmallTrendStyleProduceinfoUploadAPIResponse {
	return poolTmallTrendStyleProduceinfoUploadAPIResponse.Get().(*TmallTrendStyleProduceinfoUploadAPIResponse)
}

// ReleaseTmallTrendStyleProduceinfoUploadAPIResponse 将 TmallTrendStyleProduceinfoUploadAPIResponse 保存到 sync.Pool
func ReleaseTmallTrendStyleProduceinfoUploadAPIResponse(v *TmallTrendStyleProduceinfoUploadAPIResponse) {
	v.Reset()
	poolTmallTrendStyleProduceinfoUploadAPIResponse.Put(v)
}
