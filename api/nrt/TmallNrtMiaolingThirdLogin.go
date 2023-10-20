package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtMiaolingThirdLogin 喵零第三方免登
// tmall.nrt.miaoling.third.login
//
// 喵零第三方免登
func TmallNrtMiaolingThirdLogin(clt *core.SDKClient, req *nrt.TmallNrtMiaolingThirdLoginAPIRequest, resp *nrt.TmallNrtMiaolingThirdLoginAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
