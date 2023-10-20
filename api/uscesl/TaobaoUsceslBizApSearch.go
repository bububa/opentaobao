package uscesl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// Taobaousceslbizapsearch AP列表查询
// taobao.uscesl.biz.ap.search
//
// 查询当前门店下登记的AP列表
func Taobaousceslbizapsearch(clt *core.SDKClient, req *uscesl.TaobaousceslbizapsearchAPIRequest, session string) (*uscesl.TaobaousceslbizapsearchAPIResponse, error) {
	var resp uscesl.TaobaousceslbizapsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
