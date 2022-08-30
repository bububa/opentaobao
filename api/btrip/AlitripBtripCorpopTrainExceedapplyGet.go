package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopTrainExceedapplyGet 商旅火车票第三方超标审批单搜索接口
// alitrip.btrip.corpop.train.exceedapply.get
//
// 商旅火车票第三方超标审批单搜索接口
func AlitripBtripCorpopTrainExceedapplyGet(clt *core.SDKClient, req *btrip.AlitripBtripCorpopTrainExceedapplyGetAPIRequest, session string) (*btrip.AlitripBtripCorpopTrainExceedapplyGetAPIResponse, error) {
	var resp btrip.AlitripBtripCorpopTrainExceedapplyGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
