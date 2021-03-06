package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopMemoNoteList 查询天猫精灵用户设置的所有备忘录
// taobao.ailab.aicloud.top.memo.note.list
//
// 查询天猫精灵用户设置的所有备忘录
func TaobaoAilabAicloudTopMemoNoteList(clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMemoNoteListAPIRequest, session string) (*tmallgenie.TaobaoAilabAicloudTopMemoNoteListAPIResponse, error) {
	var resp tmallgenie.TaobaoAilabAicloudTopMemoNoteListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
