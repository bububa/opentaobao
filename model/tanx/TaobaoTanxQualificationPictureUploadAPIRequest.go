package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxQualificationPictureUploadAPIRequest
资质图片上传接口 API请求
taobao.tanx.qualification.picture.upload

资质图片上传接口 */
type TaobaoTanxQualificationPictureUploadAPIRequest struct {
	model.Params
	// dsp用户id
	_memberId int64
	// dsp用户检验token
	_token string
	// 1970年到现在的时间，毫秒
	_signTime int64
	// File文件getByte后的二进制数组
	_fileByte *model.File
}

// New
