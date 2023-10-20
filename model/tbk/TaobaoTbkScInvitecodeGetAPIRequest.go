package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscinvitecodegetAPIRequest 淘宝客-公用-私域用户邀请码生成 API请求
// taobao.tbk.sc.invitecode.get
//
// 私域用户管理(即渠道管理或会员运营管理)功能中，通过此API可生成淘宝客自身的邀请码。
type TaobaotbkscinvitecodegetAPIRequest struct {
	model.Params
	// 渠道推广的物料类型
	_relationApp string
	// 渠道关系ID
	_relationId int64
	// 邀请码类型，1 - 渠道邀请，2 - 渠道裂变，3 -会员邀请
	_codeType int64
}

// NewTaobaotbkscinvitecodegetRequest 初始化TaobaotbkscinvitecodegetAPIRequest对象
func NewTaobaotbkscinvitecodegetRequest() *TaobaotbkscinvitecodegetAPIRequest {
	return &TaobaotbkscinvitecodegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscinvitecodegetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.invitecode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscinvitecodegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscinvitecodegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRelationApp is RelationApp Setter
// 渠道推广的物料类型
func (r *TaobaotbkscinvitecodegetAPIRequest) SetRelationApp(_relationApp string) error {
	r._relationApp = _relationApp
	r.Set("relation_app", _relationApp)
	return nil
}

// GetRelationApp RelationApp Getter
func (r TaobaotbkscinvitecodegetAPIRequest) GetRelationApp() string {
	return r._relationApp
}

// SetRelationId is RelationId Setter
// 渠道关系ID
func (r *TaobaotbkscinvitecodegetAPIRequest) SetRelationId(_relationId int64) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaotbkscinvitecodegetAPIRequest) GetRelationId() int64 {
	return r._relationId
}

// SetCodeType is CodeType Setter
// 邀请码类型，1 - 渠道邀请，2 - 渠道裂变，3 -会员邀请
func (r *TaobaotbkscinvitecodegetAPIRequest) SetCodeType(_codeType int64) error {
	r._codeType = _codeType
	r.Set("code_type", _codeType)
	return nil
}

// GetCodeType CodeType Getter
func (r TaobaotbkscinvitecodegetAPIRequest) GetCodeType() int64 {
	return r._codeType
}
