package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpAccountGetCanUseBizcode 获取账户可用的bizCode
// taobao.universalbp.account.get.can.use.bizcode
//
// 查询账户可用场景，查询场景名称和场景bizcode的对应关系。其中bizcode在几乎所有接口的context中需要传入。
func TaobaoUniversalbpAccountGetCanUseBizcode(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest, resp *simba.TaobaoUniversalbpAccountGetCanUseBizcodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
