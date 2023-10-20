package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaoproductupdate 更新产品
// taobao.fenxiao.product.update
//
// 更新分销平台产品数据，不传更新数据返回失败&lt;br&gt;&lt;br/&gt;1. 对sku进行增、删操作时，原有的sku_ids字段会被忽略，请使用sku_properties和sku_properties_del。&lt;br&gt;
func Taobaofenxiaoproductupdate(clt *core.SDKClient, req *fenxiao.TaobaofenxiaoproductupdateAPIRequest, session string) (*fenxiao.TaobaofenxiaoproductupdateAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaoproductupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
