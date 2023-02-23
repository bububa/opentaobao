package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeItemTagSubmit ETC上翻商品打标接口
// alibaba.alihouse.newhome.item.tag.submit
//
// ETC上翻商品打标接口
func AlibabaAlihouseNewhomeItemTagSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeItemTagSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeItemTagSubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeItemTagSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
