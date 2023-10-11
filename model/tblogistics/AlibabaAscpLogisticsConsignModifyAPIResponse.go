package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsConsignModifyAPIResponse 修改物流公司和运单号 API返回值
// alibaba.ascp.logistics.consign.modify
//
// 修改物流公司和运单号
type AlibabaAscpLogisticsConsignModifyAPIResponse struct {
	model.CommonResponse
	AlibabaAscpLogisticsConsignModifyAPIResponseModel
}

// AlibabaAscpLogisticsConsignModifyAPIResponseModel is 修改物流公司和运单号 成功返回结果
type AlibabaAscpLogisticsConsignModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_consign_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
