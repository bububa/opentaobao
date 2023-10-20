package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingexpirepromotioncreateAPIResponse 短保优惠创建 API返回值
// alibaba.hm.marketing.expire.promotion.create
//
// 过期优惠优惠信息录入
type AlibabahmmarketingexpirepromotioncreateAPIResponse struct {
	model.CommonResponse
	AlibabahmmarketingexpirepromotioncreateAPIResponseModel
}

// AlibabahmmarketingexpirepromotioncreateAPIResponseModel is 短保优惠创建 成功返回结果
type AlibabahmmarketingexpirepromotioncreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_expire_promotion_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Datas []AlibabahmmarketingexpirepromotioncreateT `json:"datas,omitempty" xml:"datas>alibabahmmarketingexpirepromotioncreate_t,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// errorCode
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
