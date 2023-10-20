package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskucategoryupdateAPIResponse 商家类目修改接口 API返回值
// alibaba.wdk.sku.category.update
//
// 商家类目修改接口
type AlibabawdkskucategoryupdateAPIResponse struct {
	model.CommonResponse
	AlibabawdkskucategoryupdateAPIResponseModel
}

// AlibabawdkskucategoryupdateAPIResponseModel is 商家类目修改接口 成功返回结果
type AlibabawdkskucategoryupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_category_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabawdkskucategoryupdateApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
