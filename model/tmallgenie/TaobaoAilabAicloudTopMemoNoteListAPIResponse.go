package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopmemonotelistAPIResponse 查询天猫精灵用户设置的所有备忘录 API返回值
// taobao.ailab.aicloud.top.memo.note.list
//
// 查询天猫精灵用户设置的所有备忘录
type TaobaoailabaicloudtopmemonotelistAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopmemonotelistAPIResponseModel
}

// TaobaoailabaicloudtopmemonotelistAPIResponseModel is 查询天猫精灵用户设置的所有备忘录 成功返回结果
type TaobaoailabaicloudtopmemonotelistAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_memo_note_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务的结果封装
	Result *TaobaoailabaicloudtopmemonotelistResult `json:"result,omitempty" xml:"result,omitempty"`
}
