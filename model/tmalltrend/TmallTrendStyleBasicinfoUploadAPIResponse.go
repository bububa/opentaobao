package tmalltrend

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallTrendStyleBasicinfoUploadAPIResponse 3D款式基本信息同步API API返回值
// tmall.trend.style.basicinfo.upload
//
// 3D款式基本信息同步至天猫趋势中心
type TmallTrendStyleBasicinfoUploadAPIResponse struct {
	model.CommonResponse
	TmallTrendStyleBasicinfoUploadAPIResponseModel
}

// TmallTrendStyleBasicinfoUploadAPIResponseModel is 3D款式基本信息同步API 成功返回结果
type TmallTrendStyleBasicinfoUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_trend_style_basicinfo_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 款式基本信息同步处理结果描述
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 是否成功
	RespSuccess bool `json:"resp_success,omitempty" xml:"resp_success,omitempty"`
	// 错误码,1001-系统错误，1002-请求参数错误，1003-业务处理错误，1004-数据权限错误
	RespErrorCode int64 `json:"resp_error_code,omitempty" xml:"resp_error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
