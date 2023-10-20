package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpSuborderEstcontimeModifyAPIResponse 向前修改发货时效 API返回值
// alibaba.ascp.suborder.estcontime.modify
//
// 向前修改发货时效
type AlibabaAscpSuborderEstcontimeModifyAPIResponse struct {
	model.CommonResponse
	AlibabaAscpSuborderEstcontimeModifyAPIResponseModel
}

// AlibabaAscpSuborderEstcontimeModifyAPIResponseModel is 向前修改发货时效 成功返回结果
type AlibabaAscpSuborderEstcontimeModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_suborder_estcontime_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaAscpSuborderEstcontimeModifyResult `json:"result,omitempty" xml:"result,omitempty"`
}
