package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtophotwordsupdateAPIRequest 更新热词 API请求
// taobao.ailab.aicloud.top.hotwords.update
//
// 更新ASR热词
type TaobaoailabaicloudtophotwordsupdateAPIRequest struct {
	model.Params
	// schemaKey
	_schema string
	// 三方用户id
	_userId string
	// 业务类型
	_bizClass string
	// 操作类型
	_opType string
	// 热词内容
	_content *HotWordsContent
}

// NewTaobaoailabaicloudtophotwordsupdateRequest 初始化TaobaoailabaicloudtophotwordsupdateAPIRequest对象
func NewTaobaoailabaicloudtophotwordsupdateRequest() *TaobaoailabaicloudtophotwordsupdateAPIRequest {
	return &TaobaoailabaicloudtophotwordsupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtophotwordsupdateAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.hotwords.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtophotwordsupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtophotwordsupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// schemaKey
func (r *TaobaoailabaicloudtophotwordsupdateAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtophotwordsupdateAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 三方用户id
func (r *TaobaoailabaicloudtophotwordsupdateAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtophotwordsupdateAPIRequest) GetUserId() string {
	return r._userId
}

// SetBizClass is BizClass Setter
// 业务类型
func (r *TaobaoailabaicloudtophotwordsupdateAPIRequest) SetBizClass(_bizClass string) error {
	r._bizClass = _bizClass
	r.Set("biz_class", _bizClass)
	return nil
}

// GetBizClass BizClass Getter
func (r TaobaoailabaicloudtophotwordsupdateAPIRequest) GetBizClass() string {
	return r._bizClass
}

// SetOpType is OpType Setter
// 操作类型
func (r *TaobaoailabaicloudtophotwordsupdateAPIRequest) SetOpType(_opType string) error {
	r._opType = _opType
	r.Set("op_type", _opType)
	return nil
}

// GetOpType OpType Getter
func (r TaobaoailabaicloudtophotwordsupdateAPIRequest) GetOpType() string {
	return r._opType
}

// SetContent is Content Setter
// 热词内容
func (r *TaobaoailabaicloudtophotwordsupdateAPIRequest) SetContent(_content *HotWordsContent) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaoailabaicloudtophotwordsupdateAPIRequest) GetContent() *HotWordsContent {
	return r._content
}
