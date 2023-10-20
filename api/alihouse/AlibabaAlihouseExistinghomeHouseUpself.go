package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomehouseupself 房源上架
// alibaba.alihouse.existinghome.house.upself
//
// 房源信息上架
func Alibabaalihouseexistinghomehouseupself(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomehouseupselfAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomehouseupselfAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomehouseupselfAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
