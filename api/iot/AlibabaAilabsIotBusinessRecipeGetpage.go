package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// AlibabaAilabsIotBusinessRecipeGetpage 分页查询食谱
// alibaba.ailabs.iot.business.recipe.getpage
//
// 分页查询食谱数据
func AlibabaAilabsIotBusinessRecipeGetpage(ctx context.Context, clt *core.SDKClient, req *iot.AlibabaAilabsIotBusinessRecipeGetpageAPIRequest, resp *iot.AlibabaAilabsIotBusinessRecipeGetpageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
