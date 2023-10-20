package cainiaoecc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaoecc"
)

// Cainiaoeccexceptionsdelayget 菜鸟控制塔包裹滞留异常信息获取
// cainiao.ecc.exceptions.delay.get
//
// 菜鸟控制塔包裹滞留异常信息获取
func Cainiaoeccexceptionsdelayget(clt *core.SDKClient, req *cainiaoecc.CainiaoeccexceptionsdelaygetAPIRequest, session string) (*cainiaoecc.CainiaoeccexceptionsdelaygetAPIResponse, error) {
	var resp cainiaoecc.CainiaoeccexceptionsdelaygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
