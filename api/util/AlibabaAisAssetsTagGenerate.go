package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// AlibabaAisAssetsTagGenerate 基础设施资产标签生成
// alibaba.ais.assets.tag.generate
//
// 提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code生成
func AlibabaAisAssetsTagGenerate(ctx context.Context, clt *core.SDKClient, req *util.AlibabaAisAssetsTagGenerateAPIRequest, resp *util.AlibabaAisAssetsTagGenerateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
