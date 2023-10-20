package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// Alibabamosgoodssetprice 价格变更接口
// alibaba.mos.goods.setprice
//
// 价格变更接口，供供应商修改价格时使用。
func Alibabamosgoodssetprice(clt *core.SDKClient, req *moscm.AlibabamosgoodssetpriceAPIRequest, session string) (*moscm.AlibabamosgoodssetpriceAPIResponse, error) {
	var resp moscm.AlibabamosgoodssetpriceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
