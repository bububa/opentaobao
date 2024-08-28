package tttm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tttm"
)

// AliyunIndustryTttmProductsSync 天天特卖货品信息同步
// aliyun.industry.tttm.products.sync
//
// 天天特卖货品信息同步
func AliyunIndustryTttmProductsSync(ctx context.Context, clt *core.SDKClient, req *tttm.AliyunIndustryTttmProductsSyncAPIRequest, resp *tttm.AliyunIndustryTttmProductsSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
