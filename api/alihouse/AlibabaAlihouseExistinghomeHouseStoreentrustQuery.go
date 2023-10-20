package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomehousestoreentrustquery 门店委托信息查询
// alibaba.alihouse.existinghome.house.storeentrust.query
//
// 门店委托信息查询
func Alibabaalihouseexistinghomehousestoreentrustquery(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomehousestoreentrustqueryAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomehousestoreentrustqueryAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomehousestoreentrustqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
