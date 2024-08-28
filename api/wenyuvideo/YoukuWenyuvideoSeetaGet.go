package wenyuvideo

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wenyuvideo"
)

// YoukuWenyuvideoSeetaGet 只看TA
// youku.wenyuvideo.seeta.get
//
// 只看Ta对外输出
func YoukuWenyuvideoSeetaGet(ctx context.Context, clt *core.SDKClient, req *wenyuvideo.YoukuWenyuvideoSeetaGetAPIRequest, resp *wenyuvideo.YoukuWenyuvideoSeetaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
