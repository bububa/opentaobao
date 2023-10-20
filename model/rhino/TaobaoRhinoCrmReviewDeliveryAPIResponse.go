package rhino

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaorhinocrmreviewdeliveryAPIResponse crm实体预询期 API返回值
// taobao.rhino.crm.review.delivery
//
// crm实体预询期
type TaobaorhinocrmreviewdeliveryAPIResponse struct {
	model.CommonResponse
	TaobaorhinocrmreviewdeliveryAPIResponseModel
}

// TaobaorhinocrmreviewdeliveryAPIResponseModel is crm实体预询期 成功返回结果
type TaobaorhinocrmreviewdeliveryAPIResponseModel struct {
	XMLName xml.Name `xml:"rhino_crm_review_delivery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息，错误时展示
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回码，成功200/失败500
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 是否成功
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}
