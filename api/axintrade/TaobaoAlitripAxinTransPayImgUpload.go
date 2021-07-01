package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

/* TaobaoAlitripAxinTransPayImgUpload
上传图片到支付宝图片空间接口
taobao.alitrip.axin.trans.pay.img.upload

阿信供销平台-上传图片到支付宝图片空间接口 */
func TaobaoAlitripAxinTransPayImgUpload(clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransPayImgUploadAPIRequest, session string) (*axintrade.TaobaoAlitripAxinTransPayImgUploadAPIResponse, error) {
	var resp axintrade.TaobaoAlitripAxinTransPayImgUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
