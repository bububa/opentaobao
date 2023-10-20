package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopmemoalarmlistAPIResponse 天猫精灵闹钟查询 API返回值
// taobao.ailab.aicloud.top.memo.alarm.list
//
// 查询天猫精灵用户设置的所有闹钟
type TaobaoailabaicloudtopmemoalarmlistAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopmemoalarmlistAPIResponseModel
}

// TaobaoailabaicloudtopmemoalarmlistAPIResponseModel is 天猫精灵闹钟查询 成功返回结果
type TaobaoailabaicloudtopmemoalarmlistAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_memo_alarm_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务的结果封装
	Result *TaobaoailabaicloudtopmemoalarmlistResult `json:"result,omitempty" xml:"result,omitempty"`
}
