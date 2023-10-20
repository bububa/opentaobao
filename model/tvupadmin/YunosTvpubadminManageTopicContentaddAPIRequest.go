package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminmanagetopiccontentaddAPIRequest 专题新增内容 API请求
// yunos.tvpubadmin.manage.topic.contentadd
//
// 专题新增内容
type YunostvpubadminmanagetopiccontentaddAPIRequest struct {
	model.Params
	// 新增的专题内容
	_contentJson string
}

// NewYunostvpubadminmanagetopiccontentaddRequest 初始化YunostvpubadminmanagetopiccontentaddAPIRequest对象
func NewYunostvpubadminmanagetopiccontentaddRequest() *YunostvpubadminmanagetopiccontentaddAPIRequest {
	return &YunostvpubadminmanagetopiccontentaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadminmanagetopiccontentaddAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.topic.contentadd"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadminmanagetopiccontentaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadminmanagetopiccontentaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContentJson is ContentJson Setter
// 新增的专题内容
func (r *YunostvpubadminmanagetopiccontentaddAPIRequest) SetContentJson(_contentJson string) error {
	r._contentJson = _contentJson
	r.Set("content_json", _contentJson)
	return nil
}

// GetContentJson ContentJson Getter
func (r YunostvpubadminmanagetopiccontentaddAPIRequest) GetContentJson() string {
	return r._contentJson
}
