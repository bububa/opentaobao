package dengta

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dengta"
)

// AlibabaPicturesDengtaImsDouyinAccountChanged 接收发生变化的抖音帐号
// alibaba.pictures.dengta.ims.douyin.account.changed
//
// 接收发生变化的抖音帐号
func AlibabaPicturesDengtaImsDouyinAccountChanged(clt *core.SDKClient, req *dengta.AlibabaPicturesDengtaImsDouyinAccountChangedAPIRequest, resp *dengta.AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
