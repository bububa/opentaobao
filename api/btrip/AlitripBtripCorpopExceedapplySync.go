package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopExceedapplySync 第三方超标审批结果回传
// alitrip.btrip.corpop.exceedapply.sync
//
// 第三方审批单推送到企业后，企业审批结束，将审批结果回传给阿里商旅
func AlitripBtripCorpopExceedapplySync(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCorpopExceedapplySyncAPIRequest, resp *btrip.AlitripBtripCorpopExceedapplySyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
