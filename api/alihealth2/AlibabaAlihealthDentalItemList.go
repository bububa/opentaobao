package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthDentalItemList ISV获取口腔标品列表
// alibaba.alihealth.dental.item.list
//
// ISV获取口腔标品列表
func AlibabaAlihealthDentalItemList(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDentalItemListAPIRequest, resp *alihealth2.AlibabaAlihealthDentalItemListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
