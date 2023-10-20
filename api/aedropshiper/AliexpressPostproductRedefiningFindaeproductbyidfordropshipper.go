package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// Aliexpresspostproductredefiningfindaeproductbyidfordropshipper Dropshipper查找商品信息接口
// aliexpress.postproduct.redefining.findaeproductbyidfordropshipper
//
// 提供给Dropshipper的通过商品ID查找商品信息的接口，只有特定买家可以使用
func Aliexpresspostproductredefiningfindaeproductbyidfordropshipper(clt *core.SDKClient, req *aedropshiper.AliexpresspostproductredefiningfindaeproductbyidfordropshipperAPIRequest, session string) (*aedropshiper.AliexpresspostproductredefiningfindaeproductbyidfordropshipperAPIResponse, error) {
	var resp aedropshiper.AliexpresspostproductredefiningfindaeproductbyidfordropshipperAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
