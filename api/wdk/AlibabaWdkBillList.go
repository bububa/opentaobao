package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkBillList 五道口账单拉取接口
// alibaba.wdk.bill.list
//
// 五道口账单拉取接口
func AlibabaWdkBillList(clt *core.SDKClient, req *wdk.AlibabaWdkBillListAPIRequest, session string) (*wdk.AlibabaWdkBillListAPIResponse, error) {
	var resp wdk.AlibabaWdkBillListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
