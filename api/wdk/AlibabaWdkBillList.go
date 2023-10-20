package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkBillList 五道口账单拉取接口
// alibaba.wdk.bill.list
//
// 五道口账单拉取接口
func AlibabaWdkBillList(clt *core.SDKClient, req *wdk.AlibabaWdkBillListAPIRequest, resp *wdk.AlibabaWdkBillListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
