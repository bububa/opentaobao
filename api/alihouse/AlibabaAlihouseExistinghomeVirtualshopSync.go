package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeVirtualshopSync 二手房虚拟店铺数据同步
// alibaba.alihouse.existinghome.virtualshop.sync
//
// 二手房虚拟店铺数据同步
func AlibabaAlihouseExistinghomeVirtualshopSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeVirtualshopSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeVirtualshopSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
