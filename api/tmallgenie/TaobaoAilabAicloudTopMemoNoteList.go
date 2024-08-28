package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopMemoNoteList 查询天猫精灵用户设置的所有备忘录
// taobao.ailab.aicloud.top.memo.note.list
//
// 查询天猫精灵用户设置的所有备忘录
func TaobaoAilabAicloudTopMemoNoteList(ctx context.Context, clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMemoNoteListAPIRequest, resp *tmallgenie.TaobaoAilabAicloudTopMemoNoteListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
