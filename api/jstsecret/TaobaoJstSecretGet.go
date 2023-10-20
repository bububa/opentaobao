package jstsecret

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstsecret"
)

// Taobaojstsecretget 获取订单消费者的隐私号码
// taobao.jst.secret.get
//
// 根据订单号获取消费者的隐私号
func Taobaojstsecretget(clt *core.SDKClient, req *jstsecret.TaobaojstsecretgetAPIRequest, session string) (*jstsecret.TaobaojstsecretgetAPIResponse, error) {
	var resp jstsecret.TaobaojstsecretgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
