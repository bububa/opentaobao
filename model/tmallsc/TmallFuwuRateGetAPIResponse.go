package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallFuwuRateGetAPIResponse 服务商需获取到单条服务单评价信息 API返回值
// tmall.fuwu.rate.get
//
// 服务商需获取到单条服务单评价信息
type TmallFuwuRateGetAPIResponse struct {
	model.CommonResponse
	TmallFuwuRateGetAPIResponseModel
}

// TmallFuwuRateGetAPIResponseModel is 服务商需获取到单条服务单评价信息 成功返回结果
type TmallFuwuRateGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_fuwu_rate_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TmallFuwuRateGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
