package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentcoordinateuploadAPIRequest 协同单附件凭证上传 API请求
// alitrip.agent.coordinate.upload
//
// 协同单附件凭证上传
type AlitripagentcoordinateuploadAPIRequest struct {
	model.Params
	// 附件Byte[]
	_fileContent *model.File
}

// NewAlitripagentcoordinateuploadRequest 初始化AlitripagentcoordinateuploadAPIRequest对象
func NewAlitripagentcoordinateuploadRequest() *AlitripagentcoordinateuploadAPIRequest {
	return &AlitripagentcoordinateuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentcoordinateuploadAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.coordinate.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentcoordinateuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentcoordinateuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFileContent is FileContent Setter
// 附件Byte[]
func (r *AlitripagentcoordinateuploadAPIRequest) SetFileContent(_fileContent *model.File) error {
	r._fileContent = _fileContent
	r.Set("file_content", _fileContent)
	return nil
}

// GetFileContent FileContent Getter
func (r AlitripagentcoordinateuploadAPIRequest) GetFileContent() *model.File {
	return r._fileContent
}
