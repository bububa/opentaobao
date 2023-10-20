package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminmanagetopiccontenteditAPIRequest 专题关联内容编辑 API请求
// yunos.tvpubadmin.manage.topic.contentedit
//
// 编辑专题关联的内容
type YunostvpubadminmanagetopiccontenteditAPIRequest struct {
	model.Params
	// 入参，json格式
	_contentJson string
}

// NewYunostvpubadminmanagetopiccontenteditRequest 初始化YunostvpubadminmanagetopiccontenteditAPIRequest对象
func NewYunostvpubadminmanagetopiccontenteditRequest() *YunostvpubadminmanagetopiccontenteditAPIRequest {
	return &YunostvpubadminmanagetopiccontenteditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadminmanagetopiccontenteditAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.topic.contentedit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadminmanagetopiccontenteditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadminmanagetopiccontenteditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContentJson is ContentJson Setter
// 入参，json格式
func (r *YunostvpubadminmanagetopiccontenteditAPIRequest) SetContentJson(_contentJson string) error {
	r._contentJson = _contentJson
	r.Set("content_json", _contentJson)
	return nil
}

// GetContentJson ContentJson Getter
func (r YunostvpubadminmanagetopiccontenteditAPIRequest) GetContentJson() string {
	return r._contentJson
}
