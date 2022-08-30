package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseTradeQueryStatus 查询房源状态接口
// alibaba.alihouse.existinghome.house.trade.query.status
//
// 查询房源状态接口
func AlibabaAlihouseExistinghomeHouseTradeQueryStatus(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
