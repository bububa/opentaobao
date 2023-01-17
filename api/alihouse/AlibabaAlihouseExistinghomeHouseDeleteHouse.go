package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseDeleteHouse 删除房源
// alibaba.alihouse.existinghome.house.delete.house
//
// 删除房源
func AlibabaAlihouseExistinghomeHouseDeleteHouse(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseDeleteHouseAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeHouseDeleteHouseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
