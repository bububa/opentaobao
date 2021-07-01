package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

/* AlibabaFundplatformCardordersInfoQuery
根据制卡单分页查询卡信息
alibaba.fundplatform.cardorders.info.query

该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息 */
func AlibabaFundplatformCardordersInfoQuery(clt *core.SDKClient, req *fundplatform.AlibabaFundplatformCardordersInfoQueryAPIRequest, session string) (*fundplatform.AlibabaFundplatformCardordersInfoQueryAPIResponse, error) {
	var resp fundplatform.AlibabaFundplatformCardordersInfoQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
