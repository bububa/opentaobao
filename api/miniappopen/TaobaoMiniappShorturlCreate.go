package miniappopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappShorturlCreate 生成淘宝小程序短链接
// taobao.miniapp.shorturl.create
//
// 提供淘宝小程序短链接生成的能力，只允许对淘宝小程序对应的域名：https://m.duanqu.com/ 生成对应的短链接，其他域名无效
// 【特别注意：短链接有效期为30天，超过时效短链接将无法正常跳转到原始链接地址，请勿将短链接投放或装修到长期存在的入口】
func TaobaoMiniappShorturlCreate(ctx context.Context, clt *core.SDKClient, req *miniappopen.TaobaoMiniappShorturlCreateAPIRequest, resp *miniappopen.TaobaoMiniappShorturlCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
