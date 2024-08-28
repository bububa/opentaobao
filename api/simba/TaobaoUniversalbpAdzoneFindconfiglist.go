package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpAdzoneFindconfiglist 查询所有可用资源包信息
// taobao.universalbp.adzone.findconfiglist
//
// 查询该场景下，所有可用的资源包，可能存在数据重复。但是针对不同子场景和推广设置，可以选用的资源包有差异，建议关注补充文档，或者根据bp前端的限制，进行传参。
func TaobaoUniversalbpAdzoneFindconfiglist(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpAdzoneFindconfiglistAPIRequest, resp *simba.TaobaoUniversalbpAdzoneFindconfiglistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
