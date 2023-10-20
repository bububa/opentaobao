package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMemoNoteListAPIResponse 查询天猫精灵用户设置的所有备忘录 API返回值
// taobao.ailab.aicloud.top.memo.note.list
//
// 查询天猫精灵用户设置的所有备忘录
type TaobaoAilabAicloudTopMemoNoteListAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopMemoNoteListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMemoNoteListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopMemoNoteListAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopMemoNoteListAPIResponseModel is 查询天猫精灵用户设置的所有备忘录 成功返回结果
type TaobaoAilabAicloudTopMemoNoteListAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_memo_note_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务的结果封装
	Result *TaobaoAilabAicloudTopMemoNoteListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopMemoNoteListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopMemoNoteListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMemoNoteListAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopMemoNoteListAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopMemoNoteListAPIResponse
func GetTaobaoAilabAicloudTopMemoNoteListAPIResponse() *TaobaoAilabAicloudTopMemoNoteListAPIResponse {
	return poolTaobaoAilabAicloudTopMemoNoteListAPIResponse.Get().(*TaobaoAilabAicloudTopMemoNoteListAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopMemoNoteListAPIResponse 将 TaobaoAilabAicloudTopMemoNoteListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopMemoNoteListAPIResponse(v *TaobaoAilabAicloudTopMemoNoteListAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopMemoNoteListAPIResponse.Put(v)
}
