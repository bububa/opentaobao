package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtophotwordsgetAPIRequest 获取热词 API请求
// taobao.ailab.aicloud.top.hotwords.get
//
// 获取ASR热词
type TaobaoailabaicloudtophotwordsgetAPIRequest struct {
	model.Params
	// schemeKey
	_schema string
	// 三方用户id
	_userId string
	// 业务类型
	_bizClass string
}

// NewTaobaoailabaicloudtophotwordsgetRequest 初始化TaobaoailabaicloudtophotwordsgetAPIRequest对象
func NewTaobaoailabaicloudtophotwordsgetRequest() *TaobaoailabaicloudtophotwordsgetAPIRequest {
	return &TaobaoailabaicloudtophotwordsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtophotwordsgetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.hotwords.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtophotwordsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtophotwordsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// schemeKey
func (r *TaobaoailabaicloudtophotwordsgetAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtophotwordsgetAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 三方用户id
func (r *TaobaoailabaicloudtophotwordsgetAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtophotwordsgetAPIRequest) GetUserId() string {
	return r._userId
}

// SetBizClass is BizClass Setter
// 业务类型
func (r *TaobaoailabaicloudtophotwordsgetAPIRequest) SetBizClass(_bizClass string) error {
	r._bizClass = _bizClass
	r.Set("biz_class", _bizClass)
	return nil
}

// GetBizClass BizClass Getter
func (r TaobaoailabaicloudtophotwordsgetAPIRequest) GetBizClass() string {
	return r._bizClass
}
