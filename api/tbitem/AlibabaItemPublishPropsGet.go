package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// AlibabaItemPublishPropsGet 商品级联属性信息获取
// alibaba.item.publish.props.get
//
// 新商品发布，商品级联属性信息获取
func AlibabaItemPublishPropsGet(clt *core.SDKClient, req *tbitem.AlibabaItemPublishPropsGetAPIRequest, session string) (*tbitem.AlibabaItemPublishPropsGetAPIResponse, error) {
	var resp tbitem.AlibabaItemPublishPropsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
