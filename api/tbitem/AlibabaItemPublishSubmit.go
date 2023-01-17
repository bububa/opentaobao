package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// AlibabaItemPublishSubmit 商品发布
// alibaba.item.publish.submit
//
// 新商品发布，提交商品发布信息
func AlibabaItemPublishSubmit(clt *core.SDKClient, req *tbitem.AlibabaItemPublishSubmitAPIRequest, session string) (*tbitem.AlibabaItemPublishSubmitAPIResponse, error) {
	var resp tbitem.AlibabaItemPublishSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
