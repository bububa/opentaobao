package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeecodeupdate 新房货变更E码
// alibaba.alihouse.newhome.ecode.update
//
// 新房货变更E码
func Alibabaalihousenewhomeecodeupdate(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeecodeupdateAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeecodeupdateAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeecodeupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
