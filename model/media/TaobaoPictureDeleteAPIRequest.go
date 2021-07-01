package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPictureDeleteAPIRequest
删除图片空间图片 API请求
taobao.picture.delete

删除图片空间图片 */
type TaobaoPictureDeleteAPIRequest struct {
	model.Params
	// 图片ID字符串,可以一个也可以一组,用英文逗号间隔,如450,120,155.限制数量是100
	_pictureIds []string
}

// New
