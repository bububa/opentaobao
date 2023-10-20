package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseDeleteHouse 删除房源
// alibaba.alihouse.existinghome.house.delete.house
//
// 删除房源
func AlibabaAlihouseExistinghomeHouseDeleteHouse(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
