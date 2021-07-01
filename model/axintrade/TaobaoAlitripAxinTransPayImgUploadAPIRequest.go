package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripAxinTransPayImgUploadAPIRequest
上传图片到支付宝图片空间接口 API请求
taobao.alitrip.axin.trans.pay.img.upload

阿信供销平台-上传图片到支付宝图片空间接口 */
type TaobaoAlitripAxinTransPayImgUploadAPIRequest struct {
	model.Params
	// 上传图片到支付宝图片空间接口入参
	_axinPayImgUploadDTO *AxinPayImgUploadDto
	// 图片字节流
	_imgContents *model.File
}

// New
