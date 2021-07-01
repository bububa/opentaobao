package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPictureUpdateAPIRequest
修改图片名字 API请求
taobao.picture.update

修改指定图片的图片名 */
type TaobaoPictureUpdateAPIRequest struct {
	model.Params
	// 要更改名字的图片的id
	_pictureId int64
	// 新的图片名，最大长度50字符，不能为空
	_newName string
}

// New
