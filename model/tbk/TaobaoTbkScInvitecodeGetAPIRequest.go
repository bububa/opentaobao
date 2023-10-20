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
	_relationapp string
	// 渠道关系ID
	_relationid int64
	// 邀请码类型，1 - 渠道邀请，2 - 渠道裂变，3 -会员邀请
	_codetype int64
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

// SetRelationapp is Relationapp Setter
// 渠道推广的物料类型
func (r *TaobaotbkscinvitecodegetAPIRequest) SetRelationapp(_relationapp string) error {
	r._relationapp = _relationapp
	r.Set("relation_app", _relationapp)
	return nil
}

// GetRelationapp Relationapp Getter
func (r TaobaotbkscinvitecodegetAPIRequest) GetRelationapp() string {
	return r._relationapp
}

// SetRelationid is Relationid Setter
// 渠道关系ID
func (r *TaobaotbkscinvitecodegetAPIRequest) SetRelationid(_relationid int64) error {
	r._relationid = _relationid
	r.Set("relation_id", _relationid)
	return nil
}

// GetRelationid Relationid Getter
func (r TaobaotbkscinvitecodegetAPIRequest) GetRelationid() int64 {
	return r._relationid
}

// SetCodetype is Codetype Setter
// 邀请码类型，1 - 渠道邀请，2 - 渠道裂变，3 -会员邀请
func (r *TaobaotbkscinvitecodegetAPIRequest) SetCodetype(_codetype int64) error {
	r._codetype = _codetype
	r.Set("code_type", _codetype)
	return nil
}

// GetCodetype Codetype Getter
func (r TaobaotbkscinvitecodegetAPIRequest) GetCodetype() int64 {
	return r._codetype
}
