package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIResponse 码上放心数据落地-获取每天日报 API返回值
// alibaba.alihealth.drug.download.getentdailytaskdtolist
//
// 码上放心数据落地-获取每天日报
type AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIResponseModel
}

// AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIResponseModel is 码上放心数据落地-获取每天日报 成功返回结果
type AlibabaalihealthdrugdownloadgetentdailytaskdtolistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_download_getentdailytaskdtolist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的主要实体
	Model []EntDailyReportDto `json:"model,omitempty" xml:"model>ent_daily_report_dto,omitempty"`
	// Headers
	Headers string `json:"headers,omitempty" xml:"headers,omitempty"`
	// 请求回执信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 请求回执编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// http状态
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
}
