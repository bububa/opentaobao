package alink

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// AlibabaAlinkDeviceUnifyStatusSet 设置设备标准属性状态
// alibaba.alink.device.unify.status.set
//
// 操作用户绑定的设备
func AlibabaAlinkDeviceUnifyStatusSet(ctx context.Context, clt *core.SDKClient, req *alink.AlibabaAlinkDeviceUnifyStatusSetAPIRequest, resp *alink.AlibabaAlinkDeviceUnifyStatusSetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
