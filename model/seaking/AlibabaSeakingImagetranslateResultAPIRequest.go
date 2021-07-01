package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSeakingImagetranslateResultAPIRequest
获取图片翻译任务结果 API请求
alibaba.seaking.imagetranslate.result

获取图片翻译任务结果 */
type AlibabaSeakingImagetranslateResultAPIRequest struct {
	model.Params
	// token来源站点
	_tokenFrom string
	// 任务id
	_taskId int64
	// 用户token
	_token string
}

// New
