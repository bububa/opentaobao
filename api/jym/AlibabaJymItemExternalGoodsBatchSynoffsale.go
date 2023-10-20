package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// Alibabajymitemexternalgoodsbatchsynoffsale 同步下架接口
// alibaba.jym.item.external.goods.batch.synoffsale
//
// 同步下架接口
func Alibabajymitemexternalgoodsbatchsynoffsale(clt *core.SDKClient, req *jym.AlibabajymitemexternalgoodsbatchsynoffsaleAPIRequest, session string) (*jym.AlibabajymitemexternalgoodsbatchsynoffsaleAPIResponse, error) {
	var resp jym.AlibabajymitemexternalgoodsbatchsynoffsaleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
