package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniordergoodsready 备货完成
// taobao.omni.order.goods.ready
//
// 备货完成
func Taobaoomniordergoodsready(clt *core.SDKClient, req *omniorder.TaobaoomniordergoodsreadyAPIRequest, session string) (*omniorder.TaobaoomniordergoodsreadyAPIResponse, error) {
	var resp omniorder.TaobaoomniordergoodsreadyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
