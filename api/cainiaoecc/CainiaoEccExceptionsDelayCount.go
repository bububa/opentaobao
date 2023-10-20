package cainiaoecc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaoecc"
)

// Cainiaoeccexceptionsdelaycount 菜鸟控制塔包裹滞留异常统计信息获取
// cainiao.ecc.exceptions.delay.count
//
// 菜鸟控制塔包裹滞留异常统计信息获取
func Cainiaoeccexceptionsdelaycount(clt *core.SDKClient, req *cainiaoecc.CainiaoeccexceptionsdelaycountAPIRequest, session string) (*cainiaoecc.CainiaoeccexceptionsdelaycountAPIResponse, error) {
	var resp cainiaoecc.CainiaoeccexceptionsdelaycountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
