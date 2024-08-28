package nlife

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// AlibabaNlifeB2cItemDetailGet b2c码详情查询
// alibaba.nlife.b2c.item.detail.get
//
// 根据零售+平台生成的唯一码获取对应详情
func AlibabaNlifeB2cItemDetailGet(ctx context.Context, clt *core.SDKClient, req *nlife.AlibabaNlifeB2cItemDetailGetAPIRequest, resp *nlife.AlibabaNlifeB2cItemDetailGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
