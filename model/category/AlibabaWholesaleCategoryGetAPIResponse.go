package category

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWholesaleCategoryGetAPIResponse 获取类目信息 API返回值
// alibaba.wholesale.category.get
//
// 获取类目信息
type AlibabaWholesaleCategoryGetAPIResponse struct {
	model.CommonResponse
	AlibabaWholesaleCategoryGetAPIResponseModel
}

// AlibabaWholesaleCategoryGetAPIResponseModel is 获取类目信息 成功返回结果
type AlibabaWholesaleCategoryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wholesale_category_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 类目结果
	WholesaleCategoryResult *WholesaleCategoryOpenResult `json:"wholesale_category_result,omitempty" xml:"wholesale_category_result,omitempty"`
}
