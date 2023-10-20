package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaFundplatformCardordersInfoQuery 根据制卡单分页查询卡信息
// alibaba.fundplatform.cardorders.info.query
//
// 该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息
func AlibabaFundplatformCardordersInfoQuery(clt *core.SDKClient, req *fundplatform.AlibabaFundplatformCardordersInfoQueryAPIRequest, resp *fundplatform.AlibabaFundplatformCardordersInfoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
