package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemShiptimeUpdate 更新商品发货时间
// tmall.item.shiptime.update
//
// 增加更新删除商品/SKU发货时间(支持同一商品下的SKU同时批量更新)
// 1.
//
//	{
//	    &#34;shipTimeType&#34;: 2,  ----相对发货时间（值为1则为绝对发货时间）
//	    &#34;updateType&#34;: 0 ---更新SKU
//	},
//
//	按照指定SKU更新指定SKU的发货时间，如果原本是商品级发货时间，商品级发货时间也清空
//
// 2.
//
//	{
//	    &#34;shipTimeType&#34;: 0, -- 删除发货时间
//	    &#34;updateType&#34;: 0 --更新SKU
//	},
//	按照指定SKU删除指定SKU的发货时间
//
// 3.
//
//	{
//	    &#34;shipTimeType&#34;: 2,  ----相对发货时间（值为1则为绝对发货时间）
//	    &#34;updateType&#34;: 1 ---更新商品
//	},
//
//	更新商品级发货时间，如果原本是SKU级发货时间，清空所有SKU上的发货时间
//
// 4.
//
//	{
//	    &#34;shipTimeType&#34;: 0, -- 删除发货时间
//	    &#34;updateType&#34;: 1 --更新商品
//	},
//	删除商品级的发货时间
func TmallItemShiptimeUpdate(clt *core.SDKClient, req *tbitem.TmallItemShiptimeUpdateAPIRequest, session string) (*tbitem.TmallItemShiptimeUpdateAPIResponse, error) {
	var resp tbitem.TmallItemShiptimeUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
