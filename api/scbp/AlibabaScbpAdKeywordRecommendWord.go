package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordrecommendword 推词
// alibaba.scbp.ad.keyword.recommend.word
//
// 推词
func Alibabascbpadkeywordrecommendword(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordrecommendwordAPIRequest, session string) (*scbp.AlibabascbpadkeywordrecommendwordAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordrecommendwordAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
