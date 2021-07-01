package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailCommissionResultQueryAPIResponse
分佣结果查询 API返回值
alibaba.retail.commission.result.query

查询导购分佣记录 */
type AlibabaRetailCommissionResultQueryAPIResponse struct {
	model.CommonResponse
	AlibabaRetailCommissionResultQueryAPIResponseModel
}

// AlibabaRetailCommissionResultQueryAPIResponseModel is 分佣结果查询 成功返回结果
type AlibabaRetailCommissionResultQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_commission_result_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
