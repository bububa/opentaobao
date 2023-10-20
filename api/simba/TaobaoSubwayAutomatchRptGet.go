package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwayautomatchrptget 查询流量智选天级报告
// taobao.subway.automatch.rpt.get
//
// 查询流量智选天级报告
func Taobaosubwayautomatchrptget(clt *core.SDKClient, req *simba.TaobaosubwayautomatchrptgetAPIRequest, session string) (*simba.TaobaosubwayautomatchrptgetAPIResponse, error) {
	var resp simba.TaobaosubwayautomatchrptgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
