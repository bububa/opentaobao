package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// Alibabafundplatformcardordersstatusmakefinish 制卡商通知制卡完成
// alibaba.fundplatform.cardorders.status.make.finish
//
// 当制卡完成后，制卡商需要调用该接口，通知我们制卡已完成。
func Alibabafundplatformcardordersstatusmakefinish(clt *core.SDKClient, req *fundplatform.AlibabafundplatformcardordersstatusmakefinishAPIRequest, session string) (*fundplatform.AlibabafundplatformcardordersstatusmakefinishAPIResponse, error) {
	var resp fundplatform.AlibabafundplatformcardordersstatusmakefinishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
