package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordquerykeywordrankprice 查询关键词前五名排价
// alibaba.scbp.ad.keyword.query.keyword.rank.price
//
// 查询关键词前五名排价
func Alibabascbpadkeywordquerykeywordrankprice(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordquerykeywordrankpriceAPIRequest, session string) (*scbp.AlibabascbpadkeywordquerykeywordrankpriceAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordquerykeywordrankpriceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
