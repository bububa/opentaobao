package sungari

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/sungari"
)

// TaobaoSungariDisposeSubmit 商品商家处置提交任务
// taobao.sungari.dispose.submit
//
// 商品商家处置信息接口，提供政府部门发送处置信息给阿里
func TaobaoSungariDisposeSubmit(clt *core.SDKClient, req *sungari.TaobaoSungariDisposeSubmitAPIRequest, resp *sungari.TaobaoSungariDisposeSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
