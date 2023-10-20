package ma

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ma"
)

// Taobaomaqrcodecommoncreate 创建码平台常用二维码
// taobao.ma.qrcode.common.create
//
// 创建码平台对外提供的常用二维码接口，不适于码平台业务类型的码创建，如不支持包裹码、媒体码等，业务类型的码需要单独提供API。
func Taobaomaqrcodecommoncreate(clt *core.SDKClient, req *ma.TaobaomaqrcodecommoncreateAPIRequest, session string) (*ma.TaobaomaqrcodecommoncreateAPIResponse, error) {
	var resp ma.TaobaomaqrcodecommoncreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
