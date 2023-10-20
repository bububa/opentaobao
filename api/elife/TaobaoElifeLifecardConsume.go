package elife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/elife"
)

// Taobaoelifelifecardconsume 品牌惠卡券核销
// taobao.elife.lifecard.consume
//
// 用户线上购买生活汇品牌惠虚拟消费卡，线下购物时，商家码枪核销，涉及用户虚拟卡余额扣减操作
func Taobaoelifelifecardconsume(clt *core.SDKClient, req *elife.TaobaoelifelifecardconsumeAPIRequest, session string) (*elife.TaobaoelifelifecardconsumeAPIResponse, error) {
	var resp elife.TaobaoelifelifecardconsumeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
