package aedropshiper

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressdsmemberorderdatasubmitAPIResponse dropshipper数据回流 API返回值
// aliexpress.ds.member.orderdata.submit
//
// dropshipper数据回流
type AliexpressdsmemberorderdatasubmitAPIResponse struct {
	model.CommonResponse
	AliexpressdsmemberorderdatasubmitAPIResponseModel
}

// AliexpressdsmemberorderdatasubmitAPIResponseModel is dropshipper数据回流 成功返回结果
type AliexpressdsmemberorderdatasubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ds_member_orderdata_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Error message
	RspMsg string `json:"rsp_msg,omitempty" xml:"rsp_msg,omitempty"`
	// error code
	RspCode string `json:"rsp_code,omitempty" xml:"rsp_code,omitempty"`
	// Result object
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
