package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// Alibabafundplatformcardordersinfoquerybycardno 通过卡号查询卡信息
// alibaba.fundplatform.cardorders.info.query.by.cardno
//
// 该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息
func Alibabafundplatformcardordersinfoquerybycardno(clt *core.SDKClient, req *fundplatform.AlibabafundplatformcardordersinfoquerybycardnoAPIRequest, session string) (*fundplatform.AlibabafundplatformcardordersinfoquerybycardnoAPIResponse, error) {
	var resp fundplatform.AlibabafundplatformcardordersinfoquerybycardnoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
