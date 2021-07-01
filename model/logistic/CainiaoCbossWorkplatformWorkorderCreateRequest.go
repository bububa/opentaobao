package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单创建接口 API请求
cainiao.cboss.workplatform.workorder.create

菜鸟工单创建接口，目前调用者ISV
*/
type CainiaoCbossWorkplatformWorkorderCreateAPIRequest struct {
    model.Params
    // 工单类型
    _workOrderType   string
    // 业务类型
    _bizType   string
    // 工单创建备注
    _memo   string
    // 货主商家用户id
    _memberId   string
    // 货主用户角色
    _memberRole   string
    // 创建者淘宝id（区分子账号）
    _creatorId   string
    // 创建者角色
    _creatorRole   string
    // 外部业务系统主键
    _bizEntityValue   string
    // 店铺用户id
    _shopUserId   string
    // 交易订单id
    _tradeId   string
    // 工单来源
    _source   string
    // 来源签名，用于唯一区分不同的来源方
    _sourceSign   string
    // 运单号
    _mailNo   string
    // 扩展字段
    _features   string
    // 凭证地址列表
    _attachPathList   []string
}

// 初始化CainiaoCbossWorkplatformWorkorderCreateAPIRequest对象
func NewCainiaoCbossWorkplatformWorkorderCreateRequest() *CainiaoCbossWorkplatformWorkorderCreateAPIRequest{
    return &CainiaoCbossWorkplatformWorkorderCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCbossWorkplatformWorkorderCreateAPIRequest) GetApiMethodName() string {
    return "cainiao.cboss.workplatform.workorder.create"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCbossWorkplatformWorkorderCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkOrderType Setter
// 工单类型
func (r *CainiaoCbossWorkplatformWorkorderCreateAPIRequest) SetWorkOrderType(_workOrderType string) error {
    r._workOrderType = _workOrderType
    r.Set("work_order_type", _workOrderType)
    return nil
}

// WorkOrderType Getter
func (r CainiaoCbossWorkplatformWorkorderCreateAPIRequest) GetWorkOrderType() string {
    return r._workOrderType
}
// BizType Setter
// 业务类型
func (r *CainiaoCbossWorkplatformWorkorderCreateAPIRequest) SetBizType(_bizType string) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r CainiaoCbossWorkplatformWorkorderCreateAPIRequest) GetBizType() string {
    return r._bizType
}
// Memo Setter
// 工单创建备注
func (r *CainiaoCbossWorkplatformWorkorderCreateAPIRequest) SetMemo(_memo string) error {
    r._memo = _memo
    r.Set("memo", _memo)
    return nil
}

// Memo Getter
func (r CainiaoCbossWorkplatformWorkorderCreateAPIRequest) GetMemo() string {
    return r._memo
}
// MemberId Setter
// 货主商家用户id
func (r *CainiaoCbossWorkplatformWorkorderCreateAPIRequest) SetMemberId(_memberId string) error {
    r._memberId = _memberId
    r.Set("member_id", _memberId)
    return nil
}

// MemberId Getter
func (r CainiaoCbossWorkplatformWorkorderCreateAPIRequest) GetMemberId() string {
    return r._memberId
}
// MemberRole Setter
// 货主用户角色
func (r *CainiaoCbossWorkplatformWorkorderCreateAPIRequest) SetMemberRole(_memberRole string) error {
    r._memberRole = _memberRole
    r.Set("member_role", _memberRole)
    return nil
}

// MemberRole Getter
func (r CainiaoCbossWorkplatformWorkorderCreateAPIRequest) GetMemberRole() string {
    return r._memberRole
}
// CreatorId Setter
// 创建者淘宝id（区分子账号）
func (r *CainiaoCbossWorkplatformWorkorderCreateAPIRequest) SetCreatorId(_creatorId string) error {
    r._creatorId = _creatorId
    r.Set("creator_id", _creatorId)
    return nil
}

// CreatorId Getter
func (r CainiaoCbossWorkplatformWorkorderCreateAPIRequest) GetCreatorId() string {
    return r._creatorId
}
// CreatorRole Setter
// 创建者角色
func (r *CainiaoCbossWorkplatformWorkorderCreateAPIRequest) SetCreatorRole(_creatorRole string) error {
    r._creatorRole = _creatorRole
    r.Set("creator_role", _creatorRole)
    return nil
}

// CreatorRole Getter
func (r CainiaoCbossWorkplatformWorkorderCreateAPIRequest) GetCreatorRole() string {
    return r._creatorRole
}
// BizEntityValue Setter
// 外部业务系统主键
func (r *CainiaoCbossWorkplatformWorkorderCreateAPIRequest) SetBizEntityValue(_bizEntityValue string) error {
    r._bizEntityValue = _bizEntityValue
    r.Set("biz_entity_value", _bizEntityValue)
    return nil
}

// BizEntityValue Getter
func (r CainiaoCbossWorkplatformWorkorderCreateAPIRequest) GetBizEntityValue() string {
    return r._bizEntityValue
}
// ShopUserId Setter
// 店铺用户id
func (r *CainiaoCbossWorkplatformWorkorderCreateAPIRequest) SetShopUserId(_shopUserId string) error {
    r._shopUserId = _shopUserId
    r.Set("shop_user_id", _shopUserId)
    return nil
}

// ShopUserId Getter
func (r CainiaoCbossWorkplatformWorkorderCreateAPIRequest) GetShopUserId() string {
    return r._shopUserId
}
// TradeId Setter
// 交易订单id
func (r *CainiaoCbossWorkplatformWorkorderCreateAPIRequest) SetTradeId(_tradeId string) error {
    r._tradeId = _tradeId
    r.Set("trade_id", _tradeId)
    return nil
}

// TradeId Getter
func (r CainiaoCbossWorkplatformWorkorderCreateAPIRequest) GetTradeId() string {
    return r._tradeId
}
// Source Setter
// 工单来源
func (r *CainiaoCbossWorkplatformWorkorderCreateAPIRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r CainiaoCbossWorkplatformWorkorderCreateAPIRequest) GetSource() string {
    return r._source
}
// SourceSign Setter
// 来源签名，用于唯一区分不同的来源方
func (r *CainiaoCbossWorkplatformWorkorderCreateAPIRequest) SetSourceSign(_sourceSign string) error {
    r._sourceSign = _sourceSign
    r.Set("source_sign", _sourceSign)
    return nil
}

// SourceSign Getter
func (r CainiaoCbossWorkplatformWorkorderCreateAPIRequest) GetSourceSign() string {
    return r._sourceSign
}
// MailNo Setter
// 运单号
func (r *CainiaoCbossWorkplatformWorkorderCreateAPIRequest) SetMailNo(_mailNo string) error {
    r._mailNo = _mailNo
    r.Set("mail_no", _mailNo)
    return nil
}

// MailNo Getter
func (r CainiaoCbossWorkplatformWorkorderCreateAPIRequest) GetMailNo() string {
    return r._mailNo
}
// Features Setter
// 扩展字段
func (r *CainiaoCbossWorkplatformWorkorderCreateAPIRequest) SetFeatures(_features string) error {
    r._features = _features
    r.Set("features", _features)
    return nil
}

// Features Getter
func (r CainiaoCbossWorkplatformWorkorderCreateAPIRequest) GetFeatures() string {
    return r._features
}
// AttachPathList Setter
// 凭证地址列表
func (r *CainiaoCbossWorkplatformWorkorderCreateAPIRequest) SetAttachPathList(_attachPathList []string) error {
    r._attachPathList = _attachPathList
    r.Set("attach_path_list", _attachPathList)
    return nil
}

// AttachPathList Getter
func (r CainiaoCbossWorkplatformWorkorderCreateAPIRequest) GetAttachPathList() []string {
    return r._attachPathList
}
