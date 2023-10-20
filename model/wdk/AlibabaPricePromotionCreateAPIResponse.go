package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabapricepromotioncreateAPIResponse 营销档期活动创建 API返回值
// alibaba.price.promotion.create
//
// 大润发-盒马帮提供新增创建营销活动
type AlibabapricepromotioncreateAPIResponse struct {
	model.CommonResponse
	AlibabapricepromotioncreateAPIResponseModel
}

// AlibabapricepromotioncreateAPIResponseModel is 营销档期活动创建 成功返回结果
type AlibabapricepromotioncreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_price_promotion_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 错误编码，本期不作识别
	SystemCode string `json:"system_code,omitempty" xml:"system_code,omitempty"`
	// 档期活动ID
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 数量，本期不启用
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 创建是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
