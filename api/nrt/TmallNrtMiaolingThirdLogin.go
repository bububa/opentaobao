package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtmiaolingthirdlogin 喵零第三方免登
// tmall.nrt.miaoling.third.login
//
// 喵零第三方免登
func Tmallnrtmiaolingthirdlogin(clt *core.SDKClient, req *nrt.TmallnrtmiaolingthirdloginAPIRequest, session string) (*nrt.TmallnrtmiaolingthirdloginAPIResponse, error) {
	var resp nrt.TmallnrtmiaolingthirdloginAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
