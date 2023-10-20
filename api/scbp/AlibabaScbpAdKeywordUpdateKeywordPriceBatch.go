package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordupdatekeywordpricebatch 修改关键词价格
// alibaba.scbp.ad.keyword.update.keyword.price.batch
//
// 修改关键词价格
func Alibabascbpadkeywordupdatekeywordpricebatch(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordupdatekeywordpricebatchAPIRequest, session string) (*scbp.AlibabascbpadkeywordupdatekeywordpricebatchAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordupdatekeywordpricebatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
