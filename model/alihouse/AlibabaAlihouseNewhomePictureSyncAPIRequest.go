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

// NewAlibabaAlihouseNewhomePictureSyncRequest 初始化AlibabaAlihouseNewhomePictureSyncAPIRequest对象
func NewAlibabaAlihouseNewhomePictureSyncRequest() *AlibabaAlihouseNewhomePictureSyncAPIRequest {
	return &AlibabaAlihouseNewhomePictureSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomePictureSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.picture.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomePictureSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ProjectPictureData Setter
// 数据
func (r *AlibabaAlihouseNewhomePictureSyncAPIRequest) SetProjectPictureData(_projectPictureData *ProjectPictureDto) error {
	r._projectPictureData = _projectPictureData
	r.Set("project_picture_data", _projectPictureData)
	return nil
}

// Get ProjectPictureData Getter
func (r AlibabaAlihouseNewhomePictureSyncAPIRequest) GetProjectPictureData() *ProjectPictureDto {
	return r._projectPictureData
}
