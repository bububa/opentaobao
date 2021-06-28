package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单系统的工单进度下发 APIResponse
cainiao.cboss.workplatform.workorder.process.notify

菜鸟工单系统的工单进度下发（SPI）
*/
type CainiaoCbossWorkplatformWorkorderProcessNotifyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_cboss_workplatform_workorder_process_notify_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Response   *CainiaoCbossWorkplatformWorkorderProcessNotifyStruct `json:"response,omitempty" xml:"