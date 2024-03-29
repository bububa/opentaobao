package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaGspSupplyImageUploadAPIRequest gsp图片上传 API请求
// alibaba.gsp.supply.image.upload
//
// 上传图片至目标海外平台的素材空间
type AlibabaGspSupplyImageUploadAPIRequest struct {
	model.Params
	// 图片名称
	_fileName string
	// 图片文件流，像素宽度不小于500，不大于2000，像素长度不小于500，不大于2000
	_fileContent *model.File
}

// NewAlibabaGspSupplyImageUploadRequest 初始化AlibabaGspSupplyImageUploadAPIRequest对象
func NewAlibabaGspSupplyImageUploadRequest() *AlibabaGspSupplyImageUploadAPIRequest {
	return &AlibabaGspSupplyImageUploadAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaGspSupplyImageUploadAPIRequest) Reset() {
	r._fileName = ""
	r._fileContent = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaGspSupplyImageUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.gsp.supply.image.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaGspSupplyImageUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaGspSupplyImageUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFileName is FileName Setter
// 图片名称
func (r *AlibabaGspSupplyImageUploadAPIRequest) SetFileName(_fileName string) error {
	r._fileName = _fileName
	r.Set("file_name", _fileName)
	return nil
}

// GetFileName FileName Getter
func (r AlibabaGspSupplyImageUploadAPIRequest) GetFileName() string {
	return r._fileName
}

// SetFileContent is FileContent Setter
// 图片文件流，像素宽度不小于500，不大于2000，像素长度不小于500，不大于2000
func (r *AlibabaGspSupplyImageUploadAPIRequest) SetFileContent(_fileContent *model.File) error {
	r._fileContent = _fileContent
	r.Set("file_content", _fileContent)
	return nil
}

// GetFileContent FileContent Getter
func (r AlibabaGspSupplyImageUploadAPIRequest) GetFileContent() *model.File {
	return r._fileContent
}

var poolAlibabaGspSupplyImageUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaGspSupplyImageUploadRequest()
	},
}

// GetAlibabaGspSupplyImageUploadRequest 从 sync.Pool 获取 AlibabaGspSupplyImageUploadAPIRequest
func GetAlibabaGspSupplyImageUploadAPIRequest() *AlibabaGspSupplyImageUploadAPIRequest {
	return poolAlibabaGspSupplyImageUploadAPIRequest.Get().(*AlibabaGspSupplyImageUploadAPIRequest)
}

// ReleaseAlibabaGspSupplyImageUploadAPIRequest 将 AlibabaGspSupplyImageUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaGspSupplyImageUploadAPIRequest(v *AlibabaGspSupplyImageUploadAPIRequest) {
	v.Reset()
	poolAlibabaGspSupplyImageUploadAPIRequest.Put(v)
}
