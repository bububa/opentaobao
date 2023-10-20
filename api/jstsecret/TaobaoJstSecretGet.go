package jstsecret

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstsecret"
)

// TaobaoJstSecretGet 获取订单消费者的隐私号码
// taobao.jst.secret.get
//
// 根据订单号获取消费者的隐私号
func TaobaoJstSecretGet(clt *core.SDKClient, req *jstsecret.TaobaoJstSecretGetAPIRequest, resp *jstsecret.TaobaoJstSecretGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
