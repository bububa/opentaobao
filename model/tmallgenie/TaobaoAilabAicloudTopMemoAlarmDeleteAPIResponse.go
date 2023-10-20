package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMemoAlarmDeleteAPIResponse 天猫精灵闹钟删除 API返回值
// taobao.ailab.aicloud.top.memo.alarm.delete
//
// 天猫精灵闹钟删除
type TaobaoAilabAicloudTopMemoAlarmDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopMemoAlarmDeleteAPIResponseModel
}

// TaobaoAilabAicloudTopMemoAlarmDeleteAPIResponseModel is 天猫精灵闹钟删除 成功返回结果
type TaobaoAilabAicloudTopMemoAlarmDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_memo_alarm_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务的结果封装
	Result *TaobaoAilabAicloudTopMemoAlarmDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}
