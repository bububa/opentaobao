package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskuaddAPIResponse 新增商品 API返回值
// alibaba.wdk.sku.add
//
// 创建RT门店商品或DC商品
type AlibabawdkskuaddAPIResponse struct {
	model.CommonResponse
	AlibabawdkskuaddAPIResponseModel
}

// AlibabawdkskuaddAPIResponseModel is 新增商品 成功返回结果
type AlibabawdkskuaddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabawdkskuaddApiResults `json:"result,omitempty" xml:"result,omitempty"`
}
