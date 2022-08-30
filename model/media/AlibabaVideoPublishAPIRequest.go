package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaVideoPublishAPIRequest 发布视频 API请求
// alibaba.video.publish
//
// 发布视频。
// 说明：发布视频5s后再查询视频信息，否则可能无法获取播放链接
type AlibabaVideoPublishAPIRequest struct {
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

// NewAlibabaVideoPublishRequest 初始化AlibabaVideoPublishAPIRequest对象
func NewAlibabaVideoPublishRequest() *AlibabaVideoPublishAPIRequest {
	return &AlibabaVideoPublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaVideoPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.video.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaVideoPublishAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFileId is FileId Setter
// 文件id
func (r *AlibabaVideoPublishAPIRequest) SetFileId(_fileId string) error {
	r._fileId = _fileId
	r.Set("file_id", _fileId)
	return nil
}

// GetFileId FileId Getter
func (r AlibabaVideoPublishAPIRequest) GetFileId() string {
	return r._fileId
}

// SetTitle is Title Setter
// 视频标题
func (r *AlibabaVideoPublishAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r AlibabaVideoPublishAPIRequest) GetTitle() string {
	return r._title
}

// SetCoverUrl is CoverUrl Setter
// 封面图
func (r *AlibabaVideoPublishAPIRequest) SetCoverUrl(_coverUrl string) error {
	r._coverUrl = _coverUrl
	r.Set("cover_url", _coverUrl)
	return nil
}

// GetCoverUrl CoverUrl Getter
func (r AlibabaVideoPublishAPIRequest) GetCoverUrl() string {
	return r._coverUrl
}

// SetTags is Tags Setter
// 视频标签
func (r *AlibabaVideoPublishAPIRequest) SetTags(_tags string) error {
	r._tags = _tags
	r.Set("tags", _tags)
	return nil
}

// GetTags Tags Getter
func (r AlibabaVideoPublishAPIRequest) GetTags() string {
	return r._tags
}
