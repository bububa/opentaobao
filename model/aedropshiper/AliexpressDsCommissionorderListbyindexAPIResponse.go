package aedropshiper

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressDsCommissionorderListbyindexAPIResponse 联盟订单分页查询 API返回值
// aliexpress.ds.commissionorder.listbyindex
//
// 联盟订单分页查询
type AliexpressDsCommissionorderListbyindexAPIResponse struct {
	model.CommonResponse
	AliexpressDsCommissionorderListbyindexAPIResponseModel
}

// AliexpressDsCommissionorderListbyindexAPIResponseModel is 联盟订单分页查询 成功返回结果
type AliexpressDsCommissionorderListbyindexAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ds_commissionorder_listbyindex_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// System Error
	RspMsg string `json:"rsp_msg,omitempty" xml:"rsp_msg,omitempty"`
	// error code
	RspCode string `json:"rsp_code,omitempty" xml:"rsp_code,omitempty"`
	// result object
	Result *TrafficOrderResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
