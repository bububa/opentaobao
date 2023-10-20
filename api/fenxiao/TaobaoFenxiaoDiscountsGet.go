package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaodiscountsget 获取折扣信息
// taobao.fenxiao.discounts.get
//
// 查询折扣信息
func Taobaofenxiaodiscountsget(clt *core.SDKClient, req *fenxiao.TaobaofenxiaodiscountsgetAPIRequest, session string) (*fenxiao.TaobaofenxiaodiscountsgetAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaodiscountsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
