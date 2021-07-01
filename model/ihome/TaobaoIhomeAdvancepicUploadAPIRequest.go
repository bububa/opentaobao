package ihome

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoIhomeAdvancepicUploadAPIRequest
ihome图片上传 API请求
taobao.ihome.advancepic.upload

ihome 定制业务编辑器投稿素材上传 */
type TaobaoIhomeAdvancepicUploadAPIRequest struct {
	model.Params
	// 图片类
	_materials []AdvancePicMaterialDto
}

// New
