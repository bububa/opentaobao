package icbu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbucategorylevelattrgetAPIResponse 层级属性的子属性获取 API返回值
// alibaba.icbu.category.level.attr.get
//
// 用于获取层级属性（车型库）的子属性和属性值
type AlibabaicbucategorylevelattrgetAPIResponse struct {
	model.CommonResponse
	AlibabaicbucategorylevelattrgetAPIResponseModel
}

// AlibabaicbucategorylevelattrgetAPIResponseModel is 层级属性的子属性获取 成功返回结果
type AlibabaicbucategorylevelattrgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_category_level_attr_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	ResultList *AlibabaicbucategorylevelattrgetResult `json:"result_list,omitempty" xml:"result_list,omitempty"`
}
