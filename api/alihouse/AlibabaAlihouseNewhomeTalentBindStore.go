package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhometalentbindstore 达人号门店关系绑定
// alibaba.alihouse.newhome.talent.bind.store
//
// 达人号门店关系绑定
func Alibabaalihousenewhometalentbindstore(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhometalentbindstoreAPIRequest, session string) (*alihouse.AlibabaalihousenewhometalentbindstoreAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhometalentbindstoreAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
