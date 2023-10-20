package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaEleFengniaoChainstoreUpdate 修改门店信息接口
// alibaba.ele.fengniao.chainstore.update
//
// 修改门店的经纬度，文本地址，电话，门店名
func AlibabaEleFengniaoChainstoreUpdate(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoChainstoreUpdateAPIRequest, resp *logistic.AlibabaEleFengniaoChainstoreUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
