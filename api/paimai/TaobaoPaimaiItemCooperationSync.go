package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// Taobaopaimaiitemcooperationsync 商品同步
// taobao.paimai.item.cooperation.sync
//
// 商品同步
func Taobaopaimaiitemcooperationsync(clt *core.SDKClient, req *paimai.TaobaopaimaiitemcooperationsyncAPIRequest, session string) (*paimai.TaobaopaimaiitemcooperationsyncAPIResponse, error) {
	var resp paimai.TaobaopaimaiitemcooperationsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
