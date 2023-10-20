package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// Alibabapurcmallpackagesync 套餐同步
// alibaba.pur.cmall.package.sync
//
// 套餐同步
func Alibabapurcmallpackagesync(clt *core.SDKClient, req *pur.AlibabapurcmallpackagesyncAPIRequest, session string) (*pur.AlibabapurcmallpackagesyncAPIResponse, error) {
	var resp pur.AlibabapurcmallpackagesyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
