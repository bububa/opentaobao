package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// Alibabaalihealthpregnancyproductget 备孕首页获取达人配置商品
// alibaba.alihealth.pregnancy.product.get
//
// 备孕首页获取达人配置商品
func Alibabaalihealthpregnancyproductget(clt *core.SDKClient, req *alihealthcrm.AlibabaalihealthpregnancyproductgetAPIRequest, session string) (*alihealthcrm.AlibabaalihealthpregnancyproductgetAPIResponse, error) {
	var resp alihealthcrm.AlibabaalihealthpregnancyproductgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
