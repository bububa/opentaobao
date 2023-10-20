package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjocwritesaleslip 开票占库
// alibaba.mj.oc.writesaleslip
//
// 开票占库
func Alibabamjocwritesaleslip(clt *core.SDKClient, req *mos.AlibabamjocwritesaleslipAPIRequest, session string) (*mos.AlibabamjocwritesaleslipAPIResponse, error) {
	var resp mos.AlibabamjocwritesaleslipAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
