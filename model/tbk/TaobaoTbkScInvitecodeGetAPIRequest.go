package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScInvitecodeGetAPIRequest 淘宝客-公用-私域用户邀请码生成 API请求
// taobao.tbk.sc.invitecode.get
//
// 私域用户管理(即渠道管理或会员运营管理)功能中，通过此API可生成淘宝客自身的邀请码。
type TaobaoTbkScInvitecodeGetAPIRequest struct {
	model.Params
	// 渠道推广的物料类型
	_relationApp string
	// 渠道关系ID
	_relationId int64
	// 邀请码类型，1 - 渠道邀请，2 - 渠道裂变，3 -会员邀请
	_codeType int64
}

// NewTaobaoTbkScInvitecodeGetRequest 初始化TaobaoTbkScInvitecodeGetAPIRequest对象
func NewTaobaoTbkScInvitecodeGetRequest() *TaobaoTbkScInvitecodeGetAPIRequest {
	return &TaobaoTbkScInvitecodeGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkScInvitecodeGetAPIRequest) Reset() {
	r._relationApp = ""
	r._relationId = 0
	r._codeType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScInvitecodeGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.invitecode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScInvitecodeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkScInvitecodeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRelationApp is RelationApp Setter
// 渠道推广的物料类型
func (r *TaobaoTbkScInvitecodeGetAPIRequest) SetRelationApp(_relationApp string) error {
	r._relationApp = _relationApp
	r.Set("relation_app", _relationApp)
	return nil
}

// GetRelationApp RelationApp Getter
func (r TaobaoTbkScInvitecodeGetAPIRequest) GetRelationApp() string {
	return r._relationApp
}

// SetRelationId is RelationId Setter
// 渠道关系ID
func (r *TaobaoTbkScInvitecodeGetAPIRequest) SetRelationId(_relationId int64) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaoTbkScInvitecodeGetAPIRequest) GetRelationId() int64 {
	return r._relationId
}

// SetCodeType is CodeType Setter
// 邀请码类型，1 - 渠道邀请，2 - 渠道裂变，3 -会员邀请
func (r *TaobaoTbkScInvitecodeGetAPIRequest) SetCodeType(_codeType int64) error {
	r._codeType = _codeType
	r.Set("code_type", _codeType)
	return nil
}

// GetCodeType CodeType Getter
func (r TaobaoTbkScInvitecodeGetAPIRequest) GetCodeType() int64 {
	return r._codeType
}

var poolTaobaoTbkScInvitecodeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkScInvitecodeGetRequest()
	},
}

// GetTaobaoTbkScInvitecodeGetRequest 从 sync.Pool 获取 TaobaoTbkScInvitecodeGetAPIRequest
func GetTaobaoTbkScInvitecodeGetAPIRequest() *TaobaoTbkScInvitecodeGetAPIRequest {
	return poolTaobaoTbkScInvitecodeGetAPIRequest.Get().(*TaobaoTbkScInvitecodeGetAPIRequest)
}

// ReleaseTaobaoTbkScInvitecodeGetAPIRequest 将 TaobaoTbkScInvitecodeGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkScInvitecodeGetAPIRequest(v *TaobaoTbkScInvitecodeGetAPIRequest) {
	v.Reset()
	poolTaobaoTbkScInvitecodeGetAPIRequest.Put(v)
}
