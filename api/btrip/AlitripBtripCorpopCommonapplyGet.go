package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopCommonapplyGet 商旅审批单通用查询接口
// alitrip.btrip.corpop.commonapply.get
//
// 商旅审批单通用查询接口
func AlitripBtripCorpopCommonapplyGet(clt *core.SDKClient, req *btrip.AlitripBtripCorpopCommonapplyGetAPIRequest, resp *btrip.AlitripBtripCorpopCommonapplyGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
