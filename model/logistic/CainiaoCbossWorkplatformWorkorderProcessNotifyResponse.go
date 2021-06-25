package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单系统的工单进度下发 APIResponse
cainiao.cboss.workplatform.workorder.process.notify

菜鸟工单系统的工单进度下发（SPI）
*/
type CainiaoCbossWorkplatformWorkorderProcessNotifyAPIResponse struct {
    model.CommonResponse
    Response *CainiaoCbossWorkplatformWorkorderProcessNotifyResponse `json:"cainiao_cboss_workplatform_workorder_process_notify_response,omitempty"`
}

type CainiaoCbossWorkplatformWorkorderProcessNotifyResponse struct {

    // 结果
    Response   *CainiaoCbossWorkplatformWorkorderProcessNotifyStruct `json:"response,omitempty"`

}
