package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomenewrichreviewsyncAPIResponse 新评测改造接口同步 API返回值
// alibaba.alihouse.newhome.new.rich.review.sync
//
// 新评测改造接口同步
type AlibabaalihousenewhomenewrichreviewsyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomenewrichreviewsyncAPIResponseModel
}

// AlibabaalihousenewhomenewrichreviewsyncAPIResponseModel is 新评测改造接口同步 成功返回结果
type AlibabaalihousenewhomenewrichreviewsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_new_rich_review_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果实体类
	Result *AlibabaalihousenewhomenewrichreviewsyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
