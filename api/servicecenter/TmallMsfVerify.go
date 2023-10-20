package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Tmallmsfverify 喵师傅核销接口
// tmall.msf.verify
//
// msf服务核销的top接口
func Tmallmsfverify(clt *core.SDKClient, req *servicecenter.TmallmsfverifyAPIRequest, session string) (*servicecenter.TmallmsfverifyAPIResponse, error) {
	var resp servicecenter.TmallmsfverifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
