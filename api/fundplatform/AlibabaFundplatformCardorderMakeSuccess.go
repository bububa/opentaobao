package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// Alibabafundplatformcardordermakesuccess 通知制卡成功
// alibaba.fundplatform.cardorder.make.success
//
// 当外部业务方调用资金平台异步制卡接口后，资金平台制卡成功后通知异步通知业务方
func Alibabafundplatformcardordermakesuccess(clt *core.SDKClient, req *fundplatform.AlibabafundplatformcardordermakesuccessAPIRequest, session string) (*fundplatform.AlibabafundplatformcardordermakesuccessAPIResponse, error) {
	var resp fundplatform.AlibabafundplatformcardordermakesuccessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
