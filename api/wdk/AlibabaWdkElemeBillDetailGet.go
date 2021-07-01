package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkElemeBillDetailGet
饿了么对账单查询，带订单明细
alibaba.wdk.eleme.bill.detail.get

查询饿了么对账单信息，带订单明细 */
func AlibabaWdkElemeBillDetailGet(clt *core.SDKClient, req *wdk.AlibabaWdkElemeBillDetailGetAPIRequest, session string) (*wdk.AlibabaWdkElemeBillDetailGetAPIResponse, error) {
	var resp wdk.AlibabaWdkElemeBillDetailGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
