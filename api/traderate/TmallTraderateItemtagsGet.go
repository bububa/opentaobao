package traderate

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traderate"
)

// TmallTraderateItemtagsGet 通过商品ID获取标签列表
// tmall.traderate.itemtags.get
//
// 通过商品ID获取标签详细信息
func TmallTraderateItemtagsGet(clt *core.SDKClient, req *traderate.TmallTraderateItemtagsGetAPIRequest, session string) (*traderate.TmallTraderateItemtagsGetAPIResponse, error) {
	var resp traderate.TmallTraderateItemtagsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
