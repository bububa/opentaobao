package lstlogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstlogisticssendinfoqueryAPIResponse 供应商-异云-查询主订单包含的物流单 API返回值
// alibaba.lst.logistics.sendinfo.query
//
// 查询主订单包含的物流单
type AlibabalstlogisticssendinfoqueryAPIResponse struct {
	model.CommonResponse
	AlibabalstlogisticssendinfoqueryAPIResponseModel
}

// AlibabalstlogisticssendinfoqueryAPIResponseModel is 供应商-异云-查询主订单包含的物流单 成功返回结果
type AlibabalstlogisticssendinfoqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_logistics_sendinfo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabalstlogisticssendinfoqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
