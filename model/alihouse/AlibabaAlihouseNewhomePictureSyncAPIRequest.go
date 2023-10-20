package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomePictureSyncAPIRequest 图片数据同步 API请求
// alibaba.alihouse.newhome.picture.sync
//
// 图片数据同步
type AlibabaAlihouseNewhomePictureSyncAPIRequest struct {
	model.Params
	// 数据
	_projectPictureData *ProjectPictureDto
}

// NewAlibabaAlihouseNewhomePictureSyncRequest 初始化AlibabaAlihouseNewhomePictureSyncAPIRequest对象
func NewAlibabaAlihouseNewhomePictureSyncRequest() *AlibabaAlihouseNewhomePictureSyncAPIRequest {
	return &AlibabaAlihouseNewhomePictureSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomePictureSyncAPIRequest) Reset() {
	r._projectPictureData = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomePictureSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.picture.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomePictureSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomePictureSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectPictureData is ProjectPictureData Setter
// 数据
func (r *AlibabaAlihouseNewhomePictureSyncAPIRequest) SetProjectPictureData(_projectPictureData *ProjectPictureDto) error {
	r._projectPictureData = _projectPictureData
	r.Set("project_picture_data", _projectPictureData)
	return nil
}

// GetProjectPictureData ProjectPictureData Getter
func (r AlibabaAlihouseNewhomePictureSyncAPIRequest) GetProjectPictureData() *ProjectPictureDto {
	return r._projectPictureData
}

var poolAlibabaAlihouseNewhomePictureSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomePictureSyncRequest()
	},
}

// GetAlibabaAlihouseNewhomePictureSyncRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomePictureSyncAPIRequest
func GetAlibabaAlihouseNewhomePictureSyncAPIRequest() *AlibabaAlihouseNewhomePictureSyncAPIRequest {
	return poolAlibabaAlihouseNewhomePictureSyncAPIRequest.Get().(*AlibabaAlihouseNewhomePictureSyncAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomePictureSyncAPIRequest 将 AlibabaAlihouseNewhomePictureSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomePictureSyncAPIRequest(v *AlibabaAlihouseNewhomePictureSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomePictureSyncAPIRequest.Put(v)
}
