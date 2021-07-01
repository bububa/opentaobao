package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

/* AlibabaAlihealthDentalItemList
ISV获取口腔标品列表
alibaba.alihealth.dental.item.list

ISV获取口腔标品列表 */
func AlibabaAlihealthDentalItemList(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDentalItemListAPIRequest, session string) (*alihealth2.AlibabaAlihealthDentalItemListAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthDentalItemListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
