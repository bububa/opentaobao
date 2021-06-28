package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单平台根据业务类型id查询业务类型详细信息 APIResponse
cainiao.cboss.workplatform.biztype.querybyid

菜鸟工单平台根据业务类型id查询业务类型详细信息。 目前调用者ISV
*/
type CainiaoCbossWorkplatformBiztypeQuerybyidAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_cboss_workplatform_biztype_querybyid_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *CainiaoCbossWorkplatformBiztypeQuerybyidResult `json:"result,omitempty" xml:"