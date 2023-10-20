package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseSync 房源信息同步
// alibaba.alihouse.existinghome.house.sync
//
// 房源信息同步
func AlibabaAlihouseExistinghomeHouseSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeHouseSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
