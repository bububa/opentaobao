package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// AlibabaAilabsIotBusinessRecipeGetdetail 获取食谱详情
// alibaba.ailabs.iot.business.recipe.getdetail
//
// 获取食谱详情接口，获取ISV自己的食谱详情数据
func AlibabaAilabsIotBusinessRecipeGetdetail(ctx context.Context, clt *core.SDKClient, req *iot.AlibabaAilabsIotBusinessRecipeGetdetailAPIRequest, resp *iot.AlibabaAilabsIotBusinessRecipeGetdetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
