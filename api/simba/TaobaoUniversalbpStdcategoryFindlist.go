package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpStdcategoryFindlist 人群相关类目查询
// taobao.universalbp.stdcategory.findlist
//
// 入参商品信息，出参商品所属类别
func TaobaoUniversalbpStdcategoryFindlist(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpStdcategoryFindlistAPIRequest, resp *simba.TaobaoUniversalbpStdcategoryFindlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
