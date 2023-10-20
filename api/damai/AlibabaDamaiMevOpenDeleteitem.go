package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopendeleteitem 大麦换验平台-第三方对外开放-票品接口deleteItem
// alibaba.damai.mev.open.deleteitem
//
// deleteItem
func Alibabadamaimevopendeleteitem(clt *core.SDKClient, req *damai.AlibabadamaimevopendeleteitemAPIRequest, session string) (*damai.AlibabadamaimevopendeleteitemAPIResponse, error) {
	var resp damai.AlibabadamaimevopendeleteitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
