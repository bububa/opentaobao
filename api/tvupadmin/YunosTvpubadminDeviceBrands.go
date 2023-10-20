package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindevicebrands 获取终端类型下品牌列表
// yunos.tvpubadmin.device.brands
//
// 获取终端类型下品牌列表
func Yunostvpubadmindevicebrands(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindevicebrandsAPIRequest, session string) (*tvupadmin.YunostvpubadmindevicebrandsAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindevicebrandsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
