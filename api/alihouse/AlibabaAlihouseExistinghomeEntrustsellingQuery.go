package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeEntrustsellingQuery 委托信息查询接口
// alibaba.alihouse.existinghome.entrustselling.query
//
// 管家状态及房源信息接口
func AlibabaAlihouseExistinghomeEntrustsellingQuery(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeEntrustsellingQueryAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
