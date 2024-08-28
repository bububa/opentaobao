package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopMemoNoteDelete 天猫精灵备忘录删除
// taobao.ailab.aicloud.top.memo.note.delete
//
// 删除天猫精灵用户设置的备忘录
func TaobaoAilabAicloudTopMemoNoteDelete(ctx context.Context, clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest, resp *tmallgenie.TaobaoAilabAicloudTopMemoNoteDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
