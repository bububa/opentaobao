package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceBrands 获取终端类型下品牌列表
// yunos.tvpubadmin.device.brands
//
// 获取终端类型下品牌列表
func YunosTvpubadminDeviceBrands(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceBrandsAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceBrandsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
