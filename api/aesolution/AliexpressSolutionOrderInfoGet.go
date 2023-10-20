package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionorderinfoget get order detail info
// aliexpress.solution.order.info.get
//
// get order detail info
func Aliexpresssolutionorderinfoget(clt *core.SDKClient, req *aesolution.AliexpresssolutionorderinfogetAPIRequest, session string) (*aesolution.AliexpresssolutionorderinfogetAPIResponse, error) {
	var resp aesolution.AliexpresssolutionorderinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
