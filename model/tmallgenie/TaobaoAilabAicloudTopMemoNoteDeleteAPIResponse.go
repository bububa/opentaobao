package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵备忘录删除 API返回值 
taobao.ailab.aicloud.top.memo.note.delete

删除天猫精灵用户设置的备忘录
*/
type TaobaoAilabAicloudTopMemoNoteDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopMemoNoteDeleteAPIResponseModel
}

// 天猫精灵备忘录删除 成功返回结果
type TaobaoAilabAicloudTopMemoNoteDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_memo_note_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 服务的结果封装
    Result   *TaobaoAilabAicloudTopMemoNoteDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}
