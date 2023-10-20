package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeshopconfigastoresync 天猫好房店铺装修-Astore上翻
// alibaba.alihouse.newhome.shopconfig.astore.sync
//
// 天猫好房店铺装修-Astore上翻
func Alibabaalihousenewhomeshopconfigastoresync(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeshopconfigastoresyncAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeshopconfigastoresyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeshopconfigastoresyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
