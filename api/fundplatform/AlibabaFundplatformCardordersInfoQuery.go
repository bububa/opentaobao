package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// Alibabafundplatformcardordersinfoquery 根据制卡单分页查询卡信息
// alibaba.fundplatform.cardorders.info.query
//
// 该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息
func Alibabafundplatformcardordersinfoquery(clt *core.SDKClient, req *fundplatform.AlibabafundplatformcardordersinfoqueryAPIRequest, session string) (*fundplatform.AlibabafundplatformcardordersinfoqueryAPIResponse, error) {
	var resp fundplatform.AlibabafundplatformcardordersinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
