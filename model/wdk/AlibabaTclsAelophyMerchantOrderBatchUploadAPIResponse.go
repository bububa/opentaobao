package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantorderbatchuploadAPIResponse 商家订单数据批量上传 API返回值
// alibaba.tcls.aelophy.merchant.order.batch.upload
//
// 商家订单数据上传
type AlibabatclsaelophymerchantorderbatchuploadAPIResponse struct {
	model.CommonResponse
	AlibabatclsaelophymerchantorderbatchuploadAPIResponseModel
}

// AlibabatclsaelophymerchantorderbatchuploadAPIResponseModel is 商家订单数据批量上传 成功返回结果
type AlibabatclsaelophymerchantorderbatchuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_order_batch_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	ApiResult *AlibabatclsaelophymerchantorderbatchuploadApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
