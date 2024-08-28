package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopCommonapplyGet 商旅审批单通用查询接口
// alitrip.btrip.corpop.commonapply.get
//
// 商旅审批单通用查询接口
func AlitripBtripCorpopCommonapplyGet(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCorpopCommonapplyGetAPIRequest, resp *btrip.AlitripBtripCorpopCommonapplyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
