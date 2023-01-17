package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest 解除码的关联关系 API请求
// alibaba.alihealth.tracecodeseller.code.relation.codeantiactive
//
// 解除码的关联关系
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
func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopCode is TopCode Setter
// 顶层码
func (r *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest) SetTopCode(_topCode string) error {
	r._topCode = _topCode
	r.Set("top_code", _topCode)
	return nil
}

// GetTopCode TopCode Getter
func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest) GetTopCode() string {
	return r._topCode
}

// SetTbUserId is TbUserId Setter
// 淘宝名
func (r *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest) SetTbUserId(_tbUserId string) error {
	r._tbUserId = _tbUserId
	r.Set("tb_user_id", _tbUserId)
	return nil
}

// GetTbUserId TbUserId Getter
func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest) GetTbUserId() string {
	return r._tbUserId
}

// SetEntInfoId is EntInfoId Setter
// 企业id
func (r *AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest) SetEntInfoId(_entInfoId int64) error {
	r._entInfoId = _entInfoId
	r.Set("ent_info_id", _entInfoId)
	return nil
}

// GetEntInfoId EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest) GetEntInfoId() int64 {
	return r._entInfoId
}
