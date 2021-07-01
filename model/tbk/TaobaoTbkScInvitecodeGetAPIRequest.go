package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkScInvitecodeGetAPIRequest
淘宝客-公用-私域用户邀请码生成 API请求
taobao.tbk.sc.invitecode.get

私域用户管理(即渠道管理或会员运营管理)功能中，通过此API可生成淘宝客自身的邀请码。 */
type TaobaoTbkScInvitecodeGetAPIRequest struct {
	model.Params
	// 渠道关系ID
	_relationId int64
	// 渠道推广的物料类型
	_relationApp string
	// 邀请码类型，1 - 渠道邀请，2 - 渠道裂变，3 -会员邀请
	_codeType int64
}

// NewTaobaoTbkScInvitecodeGetRequest 初始化TaobaoTbkScInvitecodeGetAPIRequest对象
func NewTaobaoTbkScInvitecodeGetRequest() *TaobaoTbkScInvitecodeGetAPIRequest {
	return &TaobaoTbkScInvitecodeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScInvitecodeGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.invitecode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScInvitecodeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RelationId Setter
// 渠道关系ID
func (r *TaobaoTbkScInvitecodeGetAPIRequest) SetRelationId(_relationId int64) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// Get RelationId Getter
func (r TaobaoTbkScInvitecodeGetAPIRequest) GetRelationId() int64 {
	return r._relationId
}

// Set is RelationApp Setter
// 渠道推广的物料类型
func (r *TaobaoTbkScInvitecodeGetAPIRequest) SetRelationApp(_relationApp string) error {
	r._relationApp = _relationApp
	r.Set("relation_app", _relationApp)
	return nil
}

// Get RelationApp Getter
func (r TaobaoTbkScInvitecodeGetAPIRequest) GetRelationApp() string {
	return r._relationApp
}

// Set is CodeType Setter
// 邀请码类型，1 - 渠道邀请，2 - 渠道裂变，3 -会员邀请
func (r *TaobaoTbkScInvitecodeGetAPIRequest) SetCodeType(_codeType int64) error {
	r._codeType = _codeType
	r.Set("code_type", _codeType)
	return nil
}

// Get CodeType Getter
func (r TaobaoTbkScInvitecodeGetAPIRequest) GetCodeType() int64 {
	return r._codeType
}
