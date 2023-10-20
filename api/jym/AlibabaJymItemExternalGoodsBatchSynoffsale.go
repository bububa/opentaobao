package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymItemExternalGoodsBatchSynoffsale 同步下架接口
// alibaba.jym.item.external.goods.batch.synoffsale
//
// 同步下架接口
func AlibabaJymItemExternalGoodsBatchSynoffsale(clt *core.SDKClient, req *jym.AlibabaJymItemExternalGoodsBatchSynoffsaleAPIRequest, session string) (*jym.AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse, error) {
	var resp jym.AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
