package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabavideopublishAPIRequest 发布视频 API请求
// alibaba.video.publish
//
// 发布视频。
// 说明：发布视频5s后再查询视频信息，否则可能无法获取播放链接
type AlibabavideopublishAPIRequest struct {
	model.Params
	// 文件id
	_fileId string
	// 视频标题
	_title string
	// 封面图
	_coverUrl string
	// 视频标签
	_tags string
}

// NewAlibabavideopublishRequest 初始化AlibabavideopublishAPIRequest对象
func NewAlibabavideopublishRequest() *AlibabavideopublishAPIRequest {
	return &AlibabavideopublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabavideopublishAPIRequest) GetApiMethodName() string {
	return "alibaba.video.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabavideopublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabavideopublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFileId is FileId Setter
// 文件id
func (r *AlibabavideopublishAPIRequest) SetFileId(_fileId string) error {
	r._fileId = _fileId
	r.Set("file_id", _fileId)
	return nil
}

// GetFileId FileId Getter
func (r AlibabavideopublishAPIRequest) GetFileId() string {
	return r._fileId
}

// SetTitle is Title Setter
// 视频标题
func (r *AlibabavideopublishAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r AlibabavideopublishAPIRequest) GetTitle() string {
	return r._title
}

// SetCoverUrl is CoverUrl Setter
// 封面图
func (r *AlibabavideopublishAPIRequest) SetCoverUrl(_coverUrl string) error {
	r._coverUrl = _coverUrl
	r.Set("cover_url", _coverUrl)
	return nil
}

// GetCoverUrl CoverUrl Getter
func (r AlibabavideopublishAPIRequest) GetCoverUrl() string {
	return r._coverUrl
}

// SetTags is Tags Setter
// 视频标签
func (r *AlibabavideopublishAPIRequest) SetTags(_tags string) error {
	r._tags = _tags
	r.Set("tags", _tags)
	return nil
}

// GetTags Tags Getter
func (r AlibabavideopublishAPIRequest) GetTags() string {
	return r._tags
}
