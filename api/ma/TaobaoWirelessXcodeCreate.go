package ma

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ma"
)

// Taobaowirelessxcodecreate 创建二维码/短连接
// taobao.wireless.xcode.create
//
// 创建码平台的普通二维码或者长连接转短连接服务
func Taobaowirelessxcodecreate(clt *core.SDKClient, req *ma.TaobaowirelessxcodecreateAPIRequest, session string) (*ma.TaobaowirelessxcodecreateAPIResponse, error) {
	var resp ma.TaobaowirelessxcodecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
