package alihealthlab

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthlab"
)

// AlibabaAlihealthLabItemSync 阿里健康检验检测商品发布
// alibaba.alihealth.lab.item.sync
//
// iSV发布检验检测商品基本信息给健康，内部关联一个淘宝商品或SKU
func AlibabaAlihealthLabItemSync(ctx context.Context, clt *core.SDKClient, req *alihealthlab.AlibabaAlihealthLabItemSyncAPIRequest, resp *alihealthlab.AlibabaAlihealthLabItemSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
