package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomePictureSyncAPIRequest
图片数据同步 API请求
alibaba.alihouse.newhome.picture.sync

图片数据同步 */
type AlibabaAlihouseNewhomePictureSyncAPIRequest struct {
	model.Params
	// 数据
	_projectPictureData *ProjectPictureDto
}

// New
