package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaelophyshopupdaterangeAPIResponse 更新渠道店销售范围 API返回值
// alibaba.aelophy.shop.updaterange
//
// 更新渠道店销售范围
type AlibabaaelophyshopupdaterangeAPIResponse struct {
	model.CommonResponse
	AlibabaaelophyshopupdaterangeAPIResponseModel
}

// AlibabaaelophyshopupdaterangeAPIResponseModel is 更新渠道店销售范围 成功返回结果
type AlibabaaelophyshopupdaterangeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aelophy_shop_updaterange_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// api调用结果
	ApiResult *AlibabaaelophyshopupdaterangeApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
