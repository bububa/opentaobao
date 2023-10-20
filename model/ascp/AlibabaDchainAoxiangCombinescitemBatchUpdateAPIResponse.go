package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse 更新组合货品 API返回值
// alibaba.dchain.aoxiang.combinescitem.batch.update
//
// 更新组合货品
type AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponseModel
}

// AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponseModel is 更新组合货品 成功返回结果
type AlibabaDchainAoxiangCombinescitemBatchUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_combinescitem_batch_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	BatchUpdateCombineScitemResponse *BatchUpdateCombineScItemResponse `json:"batch_update_combine_scitem_response,omitempty" xml:"batch_update_combine_scitem_response,omitempty"`
}
