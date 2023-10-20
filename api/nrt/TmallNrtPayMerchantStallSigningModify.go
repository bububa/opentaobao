package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtpaymerchantstallsigningmodify 三级商户进件修改
// tmall.nrt.pay.merchant.stall.signing.modify
//
// 三级商户进件修改
func Tmallnrtpaymerchantstallsigningmodify(clt *core.SDKClient, req *nrt.TmallnrtpaymerchantstallsigningmodifyAPIRequest, session string) (*nrt.TmallnrtpaymerchantstallsigningmodifyAPIResponse, error) {
	var resp nrt.TmallnrtpaymerchantstallsigningmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
