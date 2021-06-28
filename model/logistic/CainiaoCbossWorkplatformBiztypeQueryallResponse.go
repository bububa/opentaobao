package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单平台根据交易订单查询某条业务线上的所有业务类型 APIResponse
cainiao.cboss.workplatform.biztype.queryall

菜鸟工单平台根据交易订单查询某条业务线上的所有业务类型。 目前调用者ISV
*/
type CainiaoCbossWorkplatformBiztypeQueryallAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_cboss_workplatform_biztype_queryall_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // bizTypeJson
    
    BizTypeJson   string `json:"biz_type_json,omitempty" xml:"