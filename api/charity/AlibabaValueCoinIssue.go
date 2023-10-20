package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabavaluecoinissue 爱豆发放
// alibaba.value.coin.issue
//
// 爱豆发放
func Alibabavaluecoinissue(clt *core.SDKClient, req *charity.AlibabavaluecoinissueAPIRequest, session string) (*charity.AlibabavaluecoinissueAPIResponse, error) {
	var resp charity.AlibabavaluecoinissueAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
