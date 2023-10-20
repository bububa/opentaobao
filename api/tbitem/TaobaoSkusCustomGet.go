package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Taobaoskuscustomget 根据外部ID取商品SKU
// taobao.skus.custom.get
//
// 跟据卖家设定的Sku的外部id获取商品，如果一个outer_id对应多个Sku会返回所有符合条件的sku &lt;br/&gt;这个Sku所属卖家从传入的session中获取，需要session绑定(注：iid标签里是num_iid的值，可以用作num_iid使用)
func Taobaoskuscustomget(clt *core.SDKClient, req *tbitem.TaobaoskuscustomgetAPIRequest, session string) (*tbitem.TaobaoskuscustomgetAPIResponse, error) {
	var resp tbitem.TaobaoskuscustomgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
