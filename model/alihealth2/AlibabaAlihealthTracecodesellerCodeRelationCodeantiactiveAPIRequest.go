package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest
解除码的关联关系 API请求
alibaba.alihealth.tracecodeseller.code.relation.codeantiactive

解除码的关联关系 */
type AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest struct {
	model.Params
	// 顶层码
	_topCode string
	// 淘宝名
	_tbUserId string
	// 企业id
	_entInfoId int64
}

// NewAlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest 初始化AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest对象
func NewAlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveRequest() *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest {
	return &AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.code.relation.codeantiactive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TopCode Setter
// 顶层码
func (r *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest) SetTopCode(_topCode string) error {
	r._topCode = _topCode
	r.Set("top_code", _topCode)
	return nil
}

// Get TopCode Getter
func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest) GetTopCode() string {
	return r._topCode
}

// Set is TbUserId Setter
// 淘宝名
func (r *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest) SetTbUserId(_tbUserId string) error {
	r._tbUserId = _tbUserId
	r.Set("tb_user_id", _tbUserId)
	return nil
}

// Get TbUserId Getter
func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest) GetTbUserId() string {
	return r._tbUserId
}

// Set is EntInfoId Setter
// 企业id
func (r *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest) SetEntInfoId(_entInfoId int64) error {
	r._entInfoId = _entInfoId
	r.Set("ent_info_id", _entInfoId)
	return nil
}

// Get EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest) GetEntInfoId() int64 {
	return r._entInfoId
}
