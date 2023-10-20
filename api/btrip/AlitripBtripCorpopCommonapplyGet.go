package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopCommonapplyGet 商旅审批单通用查询接口
// alitrip.btrip.corpop.commonapply.get
//
// 商旅审批单通用查询接口
func AlitripBtripCorpopCommonapplyGet(clt *core.SDKClient, req *btrip.AlitripBtripCorpopCommonapplyGetAPIRequest, session string) (*btrip.AlitripBtripCorpopCommonapplyGetAPIResponse, error) {
	var resp btrip.AlitripBtripCorpopCommonapplyGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
