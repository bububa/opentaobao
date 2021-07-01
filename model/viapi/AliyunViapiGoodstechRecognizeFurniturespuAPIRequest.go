package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunViapiGoodstechRecognizeFurniturespuAPIRequest
家居SPU识别 API请求
aliyun.viapi.goodstech.recognize.furniturespu

对输入的家居模型图进行分类，目前类别数可达70类。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html ) */
type AliyunViapiGoodstechRecognizeFurniturespuAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
	// 模型x方向的尺寸，单位cm，默认值100
	_xLength string
	// 模型y方向的尺寸，单位cm，默认值100
	_yLength string
	// 模型z方向的尺寸，单位cm，默认值100
	_zLength string
}

// New
