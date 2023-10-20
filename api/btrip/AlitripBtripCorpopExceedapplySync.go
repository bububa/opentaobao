package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcorpopexceedapplysync 第三方超标审批结果回传
// alitrip.btrip.corpop.exceedapply.sync
//
// 第三方审批单推送到企业后，企业审批结束，将审批结果回传给阿里商旅
func Alitripbtripcorpopexceedapplysync(clt *core.SDKClient, req *btrip.AlitripbtripcorpopexceedapplysyncAPIRequest, session string) (*btrip.AlitripbtripcorpopexceedapplysyncAPIResponse, error) {
	var resp btrip.AlitripbtripcorpopexceedapplysyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
