package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

/* AlibabaAlihealthDentalItemUnbind
ISV解绑商品
alibaba.alihealth.dental.item.unbind

ISV解绑商品 */
func AlibabaAlihealthDentalItemUnbind(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDentalItemUnbindAPIRequest, session string) (*alihealth2.AlibabaAlihealthDentalItemUnbindAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthDentalItemUnbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
