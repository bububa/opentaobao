package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabagspsupplyimageuploadAPIRequest gsp图片上传 API请求
// alibaba.gsp.supply.image.upload
//
// 上传图片至目标海外平台的素材空间
type AlibabagspsupplyimageuploadAPIRequest struct {
	model.Params
	// 图片名称
	_fileName string
	// 图片文件流，像素宽度不小于500，不大于2000，像素长度不小于500，不大于2000
	_fileContent *model.File
}

// NewAlibabagspsupplyimageuploadRequest 初始化AlibabagspsupplyimageuploadAPIRequest对象
func NewAlibabagspsupplyimageuploadRequest() *AlibabagspsupplyimageuploadAPIRequest {
	return &AlibabagspsupplyimageuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabagspsupplyimageuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.gsp.supply.image.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabagspsupplyimageuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabagspsupplyimageuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFileName is FileName Setter
// 图片名称
func (r *AlibabagspsupplyimageuploadAPIRequest) SetFileName(_fileName string) error {
	r._fileName = _fileName
	r.Set("file_name", _fileName)
	return nil
}

// GetFileName FileName Getter
func (r AlibabagspsupplyimageuploadAPIRequest) GetFileName() string {
	return r._fileName
}

// SetFileContent is FileContent Setter
// 图片文件流，像素宽度不小于500，不大于2000，像素长度不小于500，不大于2000
func (r *AlibabagspsupplyimageuploadAPIRequest) SetFileContent(_fileContent *model.File) error {
	r._fileContent = _fileContent
	r.Set("file_content", _fileContent)
	return nil
}

// GetFileContent FileContent Getter
func (r AlibabagspsupplyimageuploadAPIRequest) GetFileContent() *model.File {
	return r._fileContent
}
