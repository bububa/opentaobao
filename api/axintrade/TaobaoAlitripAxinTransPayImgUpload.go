package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripAxinTransPayImgUpload 上传图片到支付宝图片空间接口
// taobao.alitrip.axin.trans.pay.img.upload
//
// 阿信供销平台-上传图片到支付宝图片空间接口
func TaobaoAlitripAxinTransPayImgUpload(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransPayImgUploadAPIRequest, resp *axintrade.TaobaoAlitripAxinTransPayImgUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
