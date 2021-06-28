package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
dsp托管创意新增接口 APIResponse
taobao.tanx.audit.depositcreative.add

dsp托管创意新增接口
*/
type TaobaoTanxAuditDepositcreativeAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tanx_audit_depositcreative_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用的成功信息或失败信息
    
    Message   string `json:"message,omitempty" xml:"