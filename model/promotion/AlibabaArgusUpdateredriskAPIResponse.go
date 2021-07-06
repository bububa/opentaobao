package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaArgusUpdateredriskAPIResponse 更新红线价格 API返回值
// alibaba.argus.updateredrisk
//
// 商品健康中心新增红线价格规则
type AlibabaArgusUpdateredriskAPIResponse struct {
	model.CommonResponse
	AlibabaArgusUpdateredriskAPIResponseModel
}

// AlibabaArgusUpdateredriskAPIResponseModel is 更新红线价格 成功返回结果
type AlibabaArgusUpdateredriskAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_argus_updateredrisk_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 结果是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
