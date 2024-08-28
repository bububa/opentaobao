package dengta

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dengta"
)

// AlibabaPicturesDengtaImsDouyinAccountChanged 接收发生变化的抖音帐号
// alibaba.pictures.dengta.ims.douyin.account.changed
//
// 接收发生变化的抖音帐号
func AlibabaPicturesDengtaImsDouyinAccountChanged(ctx context.Context, clt *core.SDKClient, req *dengta.AlibabaPicturesDengtaImsDouyinAccountChangedAPIRequest, resp *dengta.AlibabaPicturesDengtaImsDouyinAccountChangedAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
