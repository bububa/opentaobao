package ioti

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ioti"
)

// AlibabaItAlbumDeviceSendimage 相框设备厂测刷图接口
// alibaba.it.album.device.sendimage
//
// 提供传入电子相框设备mac，mac需属于厂测白名单设备，将设备刷新为系统默认的厂测图片
func AlibabaItAlbumDeviceSendimage(ctx context.Context, clt *core.SDKClient, req *ioti.AlibabaItAlbumDeviceSendimageAPIRequest, resp *ioti.AlibabaItAlbumDeviceSendimageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
