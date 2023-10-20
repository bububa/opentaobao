package degoperation

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// Taobaodegoperationcreateqrcode 中奖生成二维码
// taobao.degoperation.createqrcode
//
// 用户中奖后，生成二维码图片链接
func Taobaodegoperationcreateqrcode(clt *core.SDKClient, req *degoperation.TaobaodegoperationcreateqrcodeAPIRequest, session string) (*degoperation.TaobaodegoperationcreateqrcodeAPIResponse, error) {
	var resp degoperation.TaobaodegoperationcreateqrcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
