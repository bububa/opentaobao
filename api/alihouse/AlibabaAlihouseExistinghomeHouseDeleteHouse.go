package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomehousedeletehouse 删除房源
// alibaba.alihouse.existinghome.house.delete.house
//
// 删除房源
func Alibabaalihouseexistinghomehousedeletehouse(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomehousedeletehouseAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomehousedeletehouseAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomehousedeletehouseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
