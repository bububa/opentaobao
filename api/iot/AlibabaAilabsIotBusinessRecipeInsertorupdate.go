package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// AlibabaAilabsIotBusinessRecipeInsertorupdate 插入和更新食谱
// alibaba.ailabs.iot.business.recipe.insertorupdate
//
// 插入和更新食谱，将isv的食谱添加到云端进行存储
func AlibabaAilabsIotBusinessRecipeInsertorupdate(ctx context.Context, clt *core.SDKClient, req *iot.AlibabaAilabsIotBusinessRecipeInsertorupdateAPIRequest, resp *iot.AlibabaAilabsIotBusinessRecipeInsertorupdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
