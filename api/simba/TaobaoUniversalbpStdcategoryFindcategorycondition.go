package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpStdcategoryFindcategorycondition 获取类目过滤条件
// taobao.universalbp.stdcategory.findcategorycondition
//
// 查询全量类目信息列表，用于进行类目兴趣人群相关定向
func TaobaoUniversalbpStdcategoryFindcategorycondition(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpStdcategoryFindcategoryconditionAPIRequest, resp *simba.TaobaoUniversalbpStdcategoryFindcategoryconditionAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
