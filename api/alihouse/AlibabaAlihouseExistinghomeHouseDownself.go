package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomehousedownself 房源信息下架
// alibaba.alihouse.existinghome.house.downself
//
// 房源信息下架
func Alibabaalihouseexistinghomehousedownself(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomehousedownselfAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomehousedownselfAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomehousedownselfAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
