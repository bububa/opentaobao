package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangconsignorderbatchqueryAPIResponse 发货单批量查询 API返回值
// alibaba.dchain.aoxiang.consignorder.batch.query
//
// 发货单批量查询
type AlibabadchainaoxiangconsignorderbatchqueryAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangconsignorderbatchqueryAPIResponseModel
}

// AlibabadchainaoxiangconsignorderbatchqueryAPIResponseModel is 发货单批量查询 成功返回结果
type AlibabadchainaoxiangconsignorderbatchqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_consignorder_batch_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	BatchQueryConsignorderResponse *BatchQueryConsignOrderResponse `json:"batch_query_consignorder_response,omitempty" xml:"batch_query_consignorder_response,omitempty"`
}
