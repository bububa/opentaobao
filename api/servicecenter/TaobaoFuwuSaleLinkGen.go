package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoFuwuSaleLinkGen 服务平台营销链接生成接口
// taobao.fuwu.sale.link.gen
//
// 服务商通过使用该接口来产生营销链接，通过把这种链接发送给商家来做自定义人群的服务营销&lt;br&gt;&lt;br/&gt;注：session是param_str这个参数串创建者生成的session，这个创建者与入参中的nick是不一致的。例如：A开发者创建了一个param_str的字符串，要为B商家生成一个营销链接，session必须是A开发者创建的session。
func TaobaoFuwuSaleLinkGen(ctx context.Context, clt *core.SDKClient, req *servicecenter.TaobaoFuwuSaleLinkGenAPIRequest, resp *servicecenter.TaobaoFuwuSaleLinkGenAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
