package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskucategorydeleteAPIResponse 商家类目删除接口 API返回值
// alibaba.wdk.sku.category.delete
//
// 商家类目删除接口
type AlibabawdkskucategorydeleteAPIResponse struct {
	model.CommonResponse
	AlibabawdkskucategorydeleteAPIResponseModel
}

// AlibabawdkskucategorydeleteAPIResponseModel is 商家类目删除接口 成功返回结果
type AlibabawdkskucategorydeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_category_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabawdkskucategorydeleteApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
