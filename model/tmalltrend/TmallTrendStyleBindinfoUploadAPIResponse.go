package tmalltrend

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallTrendStyleBindinfoUploadAPIResponse 趋势词&款式绑定信息同步API API返回值
// tmall.trend.style.bindinfo.upload
//
// 趋势词&款式(服饰行业)绑定信息同步至平台
type TmallTrendStyleBindinfoUploadAPIResponse struct {
	model.CommonResponse
	TmallTrendStyleBindinfoUploadAPIResponseModel
}

// TmallTrendStyleBindinfoUploadAPIResponseModel is 趋势词&款式绑定信息同步API 成功返回结果
type TmallTrendStyleBindinfoUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_trend_style_bindinfo_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	RespSuccess bool `json:"resp_success,omitempty" xml:"resp_success,omitempty"`
	// 错误码,1001-系统错误，1002-请求参数错误，1003-业务处理错误，1004-数据权限错误
	RespErrorCode int64 `json:"resp_error_code,omitempty" xml:"resp_error_code,omitempty"`
	// 趋势词&款式关联信息同步处理结果描述
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
