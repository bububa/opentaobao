package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBrokerSubmit 提交经纪人信息
// alibaba.alihouse.existinghome.broker.submit
//
// 提交经纪人信息
func AlibabaAlihouseExistinghomeBrokerSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeBrokerSubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeBrokerSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
