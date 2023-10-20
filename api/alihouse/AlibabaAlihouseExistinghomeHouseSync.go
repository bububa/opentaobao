package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseSync 房源信息同步
// alibaba.alihouse.existinghome.house.sync
//
// 房源信息同步
func AlibabaAlihouseExistinghomeHouseSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeHouseSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeHouseSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
