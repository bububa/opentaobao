package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPictureReplaceAPIRequest
替换图片 API请求
taobao.picture.replace

替换一张图片，只替换图片数据，图片名称，图片分类等保持不变。 */
type TaobaoPictureReplaceAPIRequest struct {
	model.Params
	// 要替换的图片的id，必须大于0
	_pictureId int64
	// 图片二进制文件流,不能为空,允许png、jpg、gif图片格式
	_imageData *model.File
}

// New
