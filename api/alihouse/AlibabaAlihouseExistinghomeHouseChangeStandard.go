package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseChangeStandard 委托房源变更标准房源
// alibaba.alihouse.existinghome.house.change.standard
//
// 委托房源变更标准房源
func AlibabaAlihouseExistinghomeHouseChangeStandard(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
