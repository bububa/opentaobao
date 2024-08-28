package seaking

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// AlibabaSeakingImagetranslateResult 获取图片翻译任务结果
// alibaba.seaking.imagetranslate.result
//
// 获取图片翻译任务结果
func AlibabaSeakingImagetranslateResult(ctx context.Context, clt *core.SDKClient, req *seaking.AlibabaSeakingImagetranslateResultAPIRequest, resp *seaking.AlibabaSeakingImagetranslateResultAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
