package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductUpdate 更新产品
// taobao.fenxiao.product.update
//
// 更新分销平台产品数据，不传更新数据返回失败&lt;br&gt;&lt;br/&gt;1. 对sku进行增、删操作时，原有的sku_ids字段会被忽略，请使用sku_properties和sku_properties_del。&lt;br&gt;
func TaobaoFenxiaoProductUpdate(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductUpdateAPIRequest, session string) (*fenxiao.TaobaoFenxiaoProductUpdateAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoProductUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
