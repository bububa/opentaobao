package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopmemoalarmdeleteAPIResponse 天猫精灵闹钟删除 API返回值
// taobao.ailab.aicloud.top.memo.alarm.delete
//
// 天猫精灵闹钟删除
type TaobaoailabaicloudtopmemoalarmdeleteAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopmemoalarmdeleteAPIResponseModel
}

// TaobaoailabaicloudtopmemoalarmdeleteAPIResponseModel is 天猫精灵闹钟删除 成功返回结果
type TaobaoailabaicloudtopmemoalarmdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_memo_alarm_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务的结果封装
	Result *TaobaoailabaicloudtopmemoalarmdeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}
