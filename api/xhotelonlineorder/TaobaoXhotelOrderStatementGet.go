package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// Taobaoxhotelorderstatementget 查询账单信息
// taobao.xhotel.order.statement.get
//
// 阿里根据此接口定义输出订单账务明细，结账状态发生变化时阿里需推送账单信息。系统商可实时调用该接口来查询订单的详情
func Taobaoxhotelorderstatementget(clt *core.SDKClient, req *xhotelonlineorder.TaobaoxhotelorderstatementgetAPIRequest, session string) (*xhotelonlineorder.TaobaoxhotelorderstatementgetAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoxhotelorderstatementgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
