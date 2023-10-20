package xhotelofficial

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelofficial"
)

// Taobaoxhotelorderofficialsettleput 官网信用住结账接口
// taobao.xhotel.order.official.settle.put
//
// 用于酒店官网信用住商家结账调用
func Taobaoxhotelorderofficialsettleput(clt *core.SDKClient, req *xhotelofficial.TaobaoxhotelorderofficialsettleputAPIRequest, session string) (*xhotelofficial.TaobaoxhotelorderofficialsettleputAPIResponse, error) {
	var resp xhotelofficial.TaobaoxhotelorderofficialsettleputAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
