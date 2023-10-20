package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordbatchquerykeywordrankprice 批量查询关键词前五名排价
// alibaba.scbp.ad.keyword.batch.query.keyword.rank.price
//
// 批量查询关键词前五名排价
func Alibabascbpadkeywordbatchquerykeywordrankprice(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordbatchquerykeywordrankpriceAPIRequest, session string) (*scbp.AlibabascbpadkeywordbatchquerykeywordrankpriceAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordbatchquerykeywordrankpriceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
