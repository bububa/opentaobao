package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Tmallcarfpcarrestpayreceive 门店线下已收尾款
// tmall.car.fpcar.restpay.receive
//
// 提供给外部(大搜或其它合作方)的接口-门店线下已收尾款(不执行分佣)
func Tmallcarfpcarrestpayreceive(clt *core.SDKClient, req *servicecenter.TmallcarfpcarrestpayreceiveAPIRequest, session string) (*servicecenter.TmallcarfpcarrestpayreceiveAPIResponse, error) {
	var resp servicecenter.TmallcarfpcarrestpayreceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
