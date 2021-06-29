package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-私域用户邀请码生成 API请求
taobao.tbk.sc.invitecode.get

私域用户管理(即渠道管理或会员运营管理)功能中，通过此API可生成淘宝客自身的邀请码。
*/
type TaobaoTbkScInvitecodeGetRequest struct {
    model.Params
    // 渠道关系ID
    relationId   int64
    // 渠道推广的物料类型
    relationApp   string
    // 邀请码类型，1 - 渠道邀请，2 - 渠道裂变，3 -会员邀请
    codeType   int64
}

// 初始化TaobaoTbkScInvitecodeGetRequest对象
func NewTaobaoTbkScInvitecodeGetRequest() *TaobaoTbkScInvitecodeGetRequest{
    return &TaobaoTbkScInvitecodeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkScInvitecodeGetRequest) GetApiMethodName() string {
    return "taobao.tbk.sc.invitecode.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkScInvitecodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RelationId Setter
// 渠道关系ID
func (r *TaobaoTbkScInvitecodeGetRequest) SetRelationId(relationId int64) error {
    r.relationId = relationId
    r.Set("relation_id", relationId)
    return nil
}

// RelationId Getter
func (r TaobaoTbkScInvitecodeGetRequest) GetRelationId() int64 {
    return r.relationId
}
// RelationApp Setter
// 渠道推广的物料类型
func (r *TaobaoTbkScInvitecodeGetRequest) SetRelationApp(relationApp string) error {
    r.relationApp = relationApp
    r.Set("relation_app", relationApp)
    return nil
}

// RelationApp Getter
func (r TaobaoTbkScInvitecodeGetRequest) GetRelationApp() string {
    return r.relationApp
}
// CodeType Setter
// 邀请码类型，1 - 渠道邀请，2 - 渠道裂变，3 -会员邀请
func (r *TaobaoTbkScInvitecodeGetRequest) SetCodeType(codeType int64) error {
    r.codeType = codeType
    r.Set("code_type", codeType)
    return nil
}

// CodeType Getter
func (r TaobaoTbkScInvitecodeGetRequest) GetCodeType() int64 {
    return r.codeType
}
