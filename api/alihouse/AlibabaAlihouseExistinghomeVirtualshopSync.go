package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomevirtualshopsync 二手房虚拟店铺数据同步
// alibaba.alihouse.existinghome.virtualshop.sync
//
// 二手房虚拟店铺数据同步
func Alibabaalihouseexistinghomevirtualshopsync(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomevirtualshopsyncAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomevirtualshopsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomevirtualshopsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
