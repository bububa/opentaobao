package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

/* AlibabaAliqinOfflineCustomerAdd
系外拉新代理商增加客户经理接口
alibaba.aliqin.offline.customer.add

阿里通信这边维护了代理商和其对应的客户经理的关系，用于业务处理，开放该接口用于代理商将他们系统下的客户经理信息同步给我们 */
func AlibabaAliqinOfflineCustomerAdd(clt *core.SDKClient, req *alicom.AlibabaAliqinOfflineCustomerAddAPIRequest, session string) (*alicom.AlibabaAliqinOfflineCustomerAddAPIResponse, error) {
	var resp alicom.AlibabaAliqinOfflineCustomerAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
