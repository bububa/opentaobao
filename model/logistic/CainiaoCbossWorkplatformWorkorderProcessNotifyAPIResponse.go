package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCbossWorkplatformWorkorderProcessNotifyAPIResponse
菜鸟工单系统的工单进度下发 API返回值
cainiao.cboss.workplatform.workorder.process.notify

菜鸟工单系统的工单进度下发（SPI） */
type CainiaoCbossWorkplatformWorkorderProcessNotifyAPIResponse struct {
	model.CommonResponse
	CainiaoCbossWorkplatformWorkorderProcessNotifyAPIResponseModel
}

// CainiaoCbossWorkplatformWorkorderProcessNotifyAPIResponseModel is 菜鸟工单系统的工单进度下发 成功返回结果
type CainiaoCbossWorkplatformWorkorderProcessNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cboss_workplatform_workorder_process_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Response *CainiaoCbossWorkplatformWorkorderProcessNotifyStruct `json:"response,omitempty" xml:"response,omitempty"`
}
