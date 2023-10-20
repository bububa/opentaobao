package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskuupdateAPIResponse 更新商品 API返回值
// alibaba.wdk.sku.update
//
// 开放商品更新接口
type AlibabawdkskuupdateAPIResponse struct {
	model.CommonResponse
	AlibabawdkskuupdateAPIResponseModel
}

// AlibabawdkskuupdateAPIResponseModel is 更新商品 成功返回结果
type AlibabawdkskuupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 执行结果
	Result *AlibabawdkskuupdateApiResults `json:"result,omitempty" xml:"result,omitempty"`
}
