package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应链中台货品修改接口 APIResponse
alibaba.ascp.cnsku.update

供应链中台货品修改接口
*/
type AlibabaAscpCnskuUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_ascp_cnsku_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 异常信息
    
    ErrorMessages   []string `json:"error_messages,omitempty" xml:"