package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomepicturesyncAPIRequest 图片数据同步 API请求
// alibaba.alihouse.newhome.picture.sync
//
// 图片数据同步
type AlibabaalihousenewhomepicturesyncAPIRequest struct {
	model.Params
	// 数据
	_projectPictureData *ProjectPictureDto
}

// NewAlibabaalihousenewhomepicturesyncRequest 初始化AlibabaalihousenewhomepicturesyncAPIRequest对象
func NewAlibabaalihousenewhomepicturesyncRequest() *AlibabaalihousenewhomepicturesyncAPIRequest {
	return &AlibabaalihousenewhomepicturesyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomepicturesyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.picture.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomepicturesyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomepicturesyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectPictureData is ProjectPictureData Setter
// 数据
func (r *AlibabaalihousenewhomepicturesyncAPIRequest) SetProjectPictureData(_projectPictureData *ProjectPictureDto) error {
	r._projectPictureData = _projectPictureData
	r.Set("project_picture_data", _projectPictureData)
	return nil
}

// GetProjectPictureData ProjectPictureData Getter
func (r AlibabaalihousenewhomepicturesyncAPIRequest) GetProjectPictureData() *ProjectPictureDto {
	return r._projectPictureData
}
