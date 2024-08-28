package oversea

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/oversea"
)

// AlibabaOverseaTranslateGet 获取文本翻译信息
// alibaba.oversea.translate.get
//
// 根据传入的文本信息，获取其目标语言的翻译结果
func AlibabaOverseaTranslateGet(ctx context.Context, clt *core.SDKClient, req *oversea.AlibabaOverseaTranslateGetAPIRequest, resp *oversea.AlibabaOverseaTranslateGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
