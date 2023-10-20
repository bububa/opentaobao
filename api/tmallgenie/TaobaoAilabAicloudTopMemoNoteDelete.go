package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Taobaoailabaicloudtopmemonotedelete 天猫精灵备忘录删除
// taobao.ailab.aicloud.top.memo.note.delete
//
// 删除天猫精灵用户设置的备忘录
func Taobaoailabaicloudtopmemonotedelete(clt *core.SDKClient, req *tmallgenie.TaobaoailabaicloudtopmemonotedeleteAPIRequest, session string) (*tmallgenie.TaobaoailabaicloudtopmemonotedeleteAPIResponse, error) {
	var resp tmallgenie.TaobaoailabaicloudtopmemonotedeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
