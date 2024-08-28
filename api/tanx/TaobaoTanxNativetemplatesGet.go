package tanx

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// TaobaoTanxNativetemplatesGet 批量获取本地模板信息
// taobao.tanx.nativetemplates.get
//
// 根据传入的本地模板ID批量返回本地模板
func TaobaoTanxNativetemplatesGet(ctx context.Context, clt *core.SDKClient, req *tanx.TaobaoTanxNativetemplatesGetAPIRequest, resp *tanx.TaobaoTanxNativetemplatesGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
