package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// Alibabapurprcreate 下pr单
// alibaba.pur.pr.create
//
// 下pr单
func Alibabapurprcreate(clt *core.SDKClient, req *pur.AlibabapurprcreateAPIRequest, session string) (*pur.AlibabapurprcreateAPIResponse, error) {
	var resp pur.AlibabapurprcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
