package elife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/elife"
)

// Taobaoelifelifecardrefund 品牌惠卡券冲正退还
// taobao.elife.lifecard.refund
//
// 淘宝生活汇消费卡虚拟卡，线下冲正退货接口
func Taobaoelifelifecardrefund(clt *core.SDKClient, req *elife.TaobaoelifelifecardrefundAPIRequest, session string) (*elife.TaobaoelifelifecardrefundAPIResponse, error) {
	var resp elife.TaobaoelifelifecardrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
