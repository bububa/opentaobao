package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkUmsReturnitemsGet 退货库位商品查询（退货出库辅助）-回流单
// alibaba.wdk.ums.returnitems.get
//
// 退货库位商品查询（退货出库辅助）-回流单
func AlibabaWdkUmsReturnitemsGet(clt *core.SDKClient, req *wdk.AlibabaWdkUmsReturnitemsGetAPIRequest, resp *wdk.AlibabaWdkUmsReturnitemsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
