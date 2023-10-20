package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbabidwordpricetools 关键词出价指导工具（新）
// taobao.simba.bidword.pricetools
//
// 关键词出价指导工具（新）
func Taobaosimbabidwordpricetools(clt *core.SDKClient, req *simba.TaobaosimbabidwordpricetoolsAPIRequest, session string) (*simba.TaobaosimbabidwordpricetoolsAPIResponse, error) {
	var resp simba.TaobaosimbabidwordpricetoolsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
