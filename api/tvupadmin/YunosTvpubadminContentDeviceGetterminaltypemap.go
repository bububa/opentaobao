package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentdevicegetterminaltypemap 获取终端类型表
// yunos.tvpubadmin.content.device.getterminaltypemap
//
// 获取终端类型表
func Yunostvpubadmincontentdevicegetterminaltypemap(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentdevicegetterminaltypemapAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentdevicegetterminaltypemapAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentdevicegetterminaltypemapAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
