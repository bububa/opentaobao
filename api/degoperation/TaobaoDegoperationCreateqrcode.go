package degoperation

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// TaobaoDegoperationCreateqrcode 中奖生成二维码
// taobao.degoperation.createqrcode
//
// 用户中奖后，生成二维码图片链接
func TaobaoDegoperationCreateqrcode(ctx context.Context, clt *core.SDKClient, req *degoperation.TaobaoDegoperationCreateqrcodeAPIRequest, resp *degoperation.TaobaoDegoperationCreateqrcodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
