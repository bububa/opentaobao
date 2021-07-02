package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductScoreGet 产品质量分查询
// alibaba.icbu.product.score.get
//
// 产品质量分查询
func AlibabaIcbuProductScoreGet(clt *core.SDKClient, req *icbu.AlibabaIcbuProductScoreGetAPIRequest, session string) (*icbu.AlibabaIcbuProductScoreGetAPIResponse, error) {
	var resp icbu.AlibabaIcbuProductScoreGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
