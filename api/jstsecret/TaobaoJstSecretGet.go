package jstsecret

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstsecret"
)

/* TaobaoJstSecretGet
获取订单消费者的隐私号码
taobao.jst.secret.get

根据订单号获取消费者的隐私号 */
func TaobaoJstSecretGet(clt *core.SDKClient, req *jstsecret.TaobaoJstSecretGetAPIRequest, session string) (*jstsecret.TaobaoJstSecretGetAPIResponse, error) {
	var resp jstsecret.TaobaoJstSecretGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
