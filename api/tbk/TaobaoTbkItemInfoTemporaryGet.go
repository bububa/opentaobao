package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkItemInfoTemporaryGet 淘宝客商品详情查询（简版）（临时接口）
// taobao.tbk.item.info.temporary.get
//
// 淘宝客商品详情查询（简版）（临时接口）
func TaobaoTbkItemInfoTemporaryGet(clt *core.SDKClient, req *tbk.TaobaoTbkItemInfoTemporaryGetAPIRequest, session string) (*tbk.TaobaoTbkItemInfoTemporaryGetAPIResponse, error) {
	var resp tbk.TaobaoTbkItemInfoTemporaryGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
