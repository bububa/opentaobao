package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtMiaolingThirdLogin 喵零第三方免登
// tmall.nrt.miaoling.third.login
//
// 喵零第三方免登
func TmallNrtMiaolingThirdLogin(clt *core.SDKClient, req *nrt.TmallNrtMiaolingThirdLoginAPIRequest, session string) (*nrt.TmallNrtMiaolingThirdLoginAPIResponse, error) {
	var resp nrt.TmallNrtMiaolingThirdLoginAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
