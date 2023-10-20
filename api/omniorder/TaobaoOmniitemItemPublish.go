package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniitemitempublish 全渠道门店商品轻发布
// taobao.omniitem.item.publish
//
// 全渠道门店商品轻发布
func Taobaoomniitemitempublish(clt *core.SDKClient, req *omniorder.TaobaoomniitemitempublishAPIRequest, session string) (*omniorder.TaobaoomniitemitempublishAPIResponse, error) {
	var resp omniorder.TaobaoomniitemitempublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
