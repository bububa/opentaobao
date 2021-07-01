package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPictureIsreferencedGetAPIRequest
图片是否被引用 API请求
taobao.picture.isreferenced.get

查询图片是否被引用，被引用返回true，未被引用返回false */
type TaobaoPictureIsreferencedGetAPIRequest struct {
	model.Params
	// 图片id
	_pictureId int64
}

// New
