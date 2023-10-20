package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbuproductscoreget 产品质量分查询
// alibaba.icbu.product.score.get
//
// 产品质量分查询
func Alibabaicbuproductscoreget(clt *core.SDKClient, req *icbu.AlibabaicbuproductscoregetAPIRequest, session string) (*icbu.AlibabaicbuproductscoregetAPIResponse, error) {
	var resp icbu.AlibabaicbuproductscoregetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
