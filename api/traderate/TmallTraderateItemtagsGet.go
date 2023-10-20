package traderate

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traderate"
)

// Tmalltraderateitemtagsget 通过商品ID获取标签列表
// tmall.traderate.itemtags.get
//
// 通过商品ID获取标签详细信息
func Tmalltraderateitemtagsget(clt *core.SDKClient, req *traderate.TmalltraderateitemtagsgetAPIRequest, session string) (*traderate.TmalltraderateitemtagsgetAPIResponse, error) {
	var resp traderate.TmalltraderateitemtagsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
