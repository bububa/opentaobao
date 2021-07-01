package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSeakingImagetranslateSubmitAPIRequest
提交图片翻译任务 API请求
alibaba.seaking.imagetranslate.submit

提交图片翻译任务 */
type AlibabaSeakingImagetranslateSubmitAPIRequest struct {
	model.Params
	// token来源站点
	_tokenFrom string
	// 子任务列表
	_imageTranslateDetailList []ImageTranslateDetailDto
	// 用户token
	_token string
}

// New
