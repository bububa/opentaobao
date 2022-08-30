package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseDownself 房源信息下架
// alibaba.alihouse.existinghome.house.downself
//
// 房源信息下架
func AlibabaAlihouseExistinghomeHouseDownself(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseDownselfAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeHouseDownselfAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeHouseDownselfAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
