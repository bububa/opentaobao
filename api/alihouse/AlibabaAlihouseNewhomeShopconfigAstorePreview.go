package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeshopconfigastorepreview 天猫好房店铺装修-地址预览
// alibaba.alihouse.newhome.shopconfig.astore.preview
//
// 天猫好房店铺装修-Astore上翻
func Alibabaalihousenewhomeshopconfigastorepreview(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeshopconfigastorepreviewAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeshopconfigastorepreviewAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeshopconfigastorepreviewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
