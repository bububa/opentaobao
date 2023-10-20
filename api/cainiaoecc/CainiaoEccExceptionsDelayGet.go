package cainiaoecc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaoecc"
)

// CainiaoEccExceptionsDelayGet 菜鸟控制塔包裹滞留异常信息获取
// cainiao.ecc.exceptions.delay.get
//
// 菜鸟控制塔包裹滞留异常信息获取
func CainiaoEccExceptionsDelayGet(clt *core.SDKClient, req *cainiaoecc.CainiaoEccExceptionsDelayGetAPIRequest, resp *cainiaoecc.CainiaoEccExceptionsDelayGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
