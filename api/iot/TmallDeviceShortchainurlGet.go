package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

/* TmallDeviceShortchainurlGet
获取二维码短链接
tmall.device.shortchainurl.get

获取二维码短链接 */
func TmallDeviceShortchainurlGet(clt *core.SDKClient, req *iot.TmallDeviceShortchainurlGetAPIRequest, session string) (*iot.TmallDeviceShortchainurlGetAPIResponse, error) {
	var resp iot.TmallDeviceShortchainurlGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
