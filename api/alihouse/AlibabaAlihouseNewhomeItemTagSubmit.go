package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeitemtagsubmit ETC上翻商品打标接口
// alibaba.alihouse.newhome.item.tag.submit
//
// ETC上翻商品打标接口
func Alibabaalihousenewhomeitemtagsubmit(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeitemtagsubmitAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeitemtagsubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeitemtagsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
