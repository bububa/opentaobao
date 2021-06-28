package tmallsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商工人客诉数据上传 APIResponse
alibaba.dchain.miaoshifu.customer.complaints.put

数字服务供应链平台提供给服务商上传工人客诉数据
*/
type AlibabaDchainMiaoshifuCustomerComplaintsPutAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_dchain_miaoshifu_customer_complaints_put_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 对外展示的错误信息
    
    DisplayMsg   string `json:"display_msg,omitempty" xml:"