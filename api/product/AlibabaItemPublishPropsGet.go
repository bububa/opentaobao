package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaItemPublishPropsGet 商品级联属性信息获取
// alibaba.item.publish.props.get
//
// 新商品发布，商品级联属性信息获取
func AlibabaItemPublishPropsGet(clt *core.SDKClient, req *product.AlibabaItemPublishPropsGetAPIRequest, session string) (*product.AlibabaItemPublishPropsGetAPIResponse, error) {
	var resp product.AlibabaItemPublishPropsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
