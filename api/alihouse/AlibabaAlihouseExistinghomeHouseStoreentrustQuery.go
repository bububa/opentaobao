package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseStoreentrustQuery 门店委托信息查询
// alibaba.alihouse.existinghome.house.storeentrust.query
//
// 门店委托信息查询
func AlibabaAlihouseExistinghomeHouseStoreentrustQuery(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseStoreentrustQueryAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeHouseStoreentrustQueryAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeHouseStoreentrustQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
