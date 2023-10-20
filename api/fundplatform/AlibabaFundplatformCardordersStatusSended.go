package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// Alibabafundplatformcardordersstatussended 制卡商通知实体卡发货完成
// alibaba.fundplatform.cardorders.status.sended
//
// 当制卡商将实体卡发货完成后，需要调用该接口，通知我们已发货。
func Alibabafundplatformcardordersstatussended(clt *core.SDKClient, req *fundplatform.AlibabafundplatformcardordersstatussendedAPIRequest, session string) (*fundplatform.AlibabafundplatformcardordersstatussendedAPIResponse, error) {
	var resp fundplatform.AlibabafundplatformcardordersstatussendedAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
