package icbu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbucategorygetnewAPIResponse （新）ICBU类目树获取接口 API返回值
// alibaba.icbu.category.get.new
//
// 获取商品发布类目
type AlibabaicbucategorygetnewAPIResponse struct {
	model.CommonResponse
	AlibabaicbucategorygetnewAPIResponseModel
}

// AlibabaicbucategorygetnewAPIResponseModel is （新）ICBU类目树获取接口 成功返回结果
type AlibabaicbucategorygetnewAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_category_get_new_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 类目信息
	Category *PostCategory `json:"category,omitempty" xml:"category,omitempty"`
}
