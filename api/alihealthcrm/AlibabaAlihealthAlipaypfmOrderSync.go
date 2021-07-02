package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthAlipaypfmOrderSync 订单数据回传接口
// alibaba.alihealth.alipaypfm.order.sync
//
// 订单数据回传接口，各个isv通过我们渠道产生订单需要回传进行统计
func AlibabaAlihealthAlipaypfmOrderSync(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthAlipaypfmOrderSyncAPIRequest, session string) (*alihealthcrm.AlibabaAlihealthAlipaypfmOrderSyncAPIResponse, error) {
	var resp alihealthcrm.AlibabaAlihealthAlipaypfmOrderSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
