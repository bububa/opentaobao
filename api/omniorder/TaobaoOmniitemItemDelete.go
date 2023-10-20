package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniitemitemdelete 全渠道商品删除
// taobao.omniitem.item.delete
//
// 全渠道商品删除，能够对门店商品库商品进行删除动作
func Taobaoomniitemitemdelete(clt *core.SDKClient, req *omniorder.TaobaoomniitemitemdeleteAPIRequest, session string) (*omniorder.TaobaoomniitemitemdeleteAPIResponse, error) {
	var resp omniorder.TaobaoomniitemitemdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
