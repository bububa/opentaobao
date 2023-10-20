package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantorderuploadAPIResponse 商家订单数据上传 API返回值
// alibaba.tcls.aelophy.merchant.order.upload
//
// 商家订单数据上传
type AlibabatclsaelophymerchantorderuploadAPIResponse struct {
	model.CommonResponse
	AlibabatclsaelophymerchantorderuploadAPIResponseModel
}

// AlibabatclsaelophymerchantorderuploadAPIResponseModel is 商家订单数据上传 成功返回结果
type AlibabatclsaelophymerchantorderuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_order_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	ApiResult *AlibabatclsaelophymerchantorderuploadApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
