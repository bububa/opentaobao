package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeEntrustsellingUpdate 管家状态及房源信息接口
// alibaba.alihouse.existinghome.entrustselling.update
//
// 管家状态及房源信息接口
func AlibabaAlihouseExistinghomeEntrustsellingUpdate(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
