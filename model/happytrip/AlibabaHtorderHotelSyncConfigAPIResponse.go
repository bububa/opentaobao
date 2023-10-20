package happytrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahtorderhotelsyncconfigAPIResponse 同步配置信息 API返回值
// alibaba.htorder.hotel.sync.config
//
// 同步配置信息
type AlibabahtorderhotelsyncconfigAPIResponse struct {
	model.CommonResponse
	AlibabahtorderhotelsyncconfigAPIResponseModel
}

// AlibabahtorderhotelsyncconfigAPIResponseModel is 同步配置信息 成功返回结果
type AlibabahtorderhotelsyncconfigAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_htorder_hotel_sync_config_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ErrNo string `json:"err_no,omitempty" xml:"err_no,omitempty"`
	// 栈信息
	StackTrace string `json:"stack_trace,omitempty" xml:"stack_trace,omitempty"`
	// 错误信息
	ErrInfo string `json:"err_info,omitempty" xml:"err_info,omitempty"`
	// 执行结果
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
	// 成功OR失败
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}
