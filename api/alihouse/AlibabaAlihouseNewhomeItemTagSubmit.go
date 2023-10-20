package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeItemTagSubmit ETC上翻商品打标接口
// alibaba.alihouse.newhome.item.tag.submit
//
// ETC上翻商品打标接口
func AlibabaAlihouseNewhomeItemTagSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeItemTagSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeItemTagSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
