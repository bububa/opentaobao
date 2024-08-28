package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeTalentBindStore 达人号门店关系绑定
// alibaba.alihouse.newhome.talent.bind.store
//
// 达人号门店关系绑定
func AlibabaAlihouseNewhomeTalentBindStore(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeTalentBindStoreAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeTalentBindStoreAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
