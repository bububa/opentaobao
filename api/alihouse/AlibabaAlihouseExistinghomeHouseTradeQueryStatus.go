package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseTradeQueryStatus 查询房源状态接口
// alibaba.alihouse.existinghome.house.trade.query.status
//
// 查询房源状态接口
func AlibabaAlihouseExistinghomeHouseTradeQueryStatus(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
