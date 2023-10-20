package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpuopselfsupplierwaybillqueryAPIResponse 商家仓自营配电子面单取号 API返回值
// alibaba.ascp.uop.self.supplier.waybill.query
//
// ERP调用打印面单取号接口
type AlibabaascpuopselfsupplierwaybillqueryAPIResponse struct {
	model.CommonResponse
	AlibabaascpuopselfsupplierwaybillqueryAPIResponseModel
}

// AlibabaascpuopselfsupplierwaybillqueryAPIResponseModel is 商家仓自营配电子面单取号 成功返回结果
type AlibabaascpuopselfsupplierwaybillqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_uop_self_supplier_waybill_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	WaybillQueryResponse *ResultWrapper `json:"waybill_query_response,omitempty" xml:"waybill_query_response,omitempty"`
}
