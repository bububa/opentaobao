package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseUpself 房源上架
// alibaba.alihouse.existinghome.house.upself
//
// 房源信息上架
func AlibabaAlihouseExistinghomeHouseUpself(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseUpselfAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeHouseUpselfAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeHouseUpselfAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
