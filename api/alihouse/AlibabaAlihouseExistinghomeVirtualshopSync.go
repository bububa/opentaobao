package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeVirtualshopSync 二手房虚拟店铺数据同步
// alibaba.alihouse.existinghome.virtualshop.sync
//
// 二手房虚拟店铺数据同步
func AlibabaAlihouseExistinghomeVirtualshopSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeVirtualshopSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
