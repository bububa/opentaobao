package ioti

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ioti"
)

// Alibabaitalbumdevicesendimage 相框设备厂测刷图接口
// alibaba.it.album.device.sendimage
//
// 提供传入电子相框设备mac，mac需属于厂测白名单设备，将设备刷新为系统默认的厂测图片
func Alibabaitalbumdevicesendimage(clt *core.SDKClient, req *ioti.AlibabaitalbumdevicesendimageAPIRequest, session string) (*ioti.AlibabaitalbumdevicesendimageAPIResponse, error) {
	var resp ioti.AlibabaitalbumdevicesendimageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
