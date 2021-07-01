package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

/* AlibabaAlicomOrderPreauthorizeCreate
业务办理结果
alibaba.alicom.order.preauthorize.create

授授权:签约结果通知 */
func AlibabaAlicomOrderPreauthorizeCreate(clt *core.SDKClient, req *alicom.AlibabaAlicomOrderPreauthorizeCreateAPIRequest, session string) (*alicom.AlibabaAlicomOrderPreauthorizeCreateAPIResponse, error) {
	var resp alicom.AlibabaAlicomOrderPreauthorizeCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
