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
type CainiaoCbossWorkplatformWorkorderCreateRequest struct {
    model.Params
    // 工单类型
    workOrderType   string
    // 业务类型
    bizType   string
    // 工单创建备注
    memo   string
    // 货主商家用户id
    memberId   string
    // 货主用户角色
    memberRole   string
    // 创建者淘宝id（区分子账号）
    creatorId   string
    // 创建者角色
    creatorRole   string
    // 外部业务系统主键
    bizEntityValue   string
    // 店铺用户id
    shopUserId   string
    // 交易订单id
    tradeId   string
    // 工单来源
    source   string
    // 来源签名，用于唯一区分不同的来源方
    sourceSign   string
    // 运单号
    mailNo   string
    // 扩展字段
    features   string
    // 凭证地址列表
    attachPathList   []string
}

// 初始化CainiaoCbossWorkplatformWorkorderCreateRequest对象
func NewCainiaoCbossWorkplatformWorkorderCreateRequest() *CainiaoCbossWorkplatformWorkorderCreateRequest{
    return &CainiaoCbossWorkplatformWorkorderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetApiMethodName() string {
    return "cainiao.cboss.workplatform.workorder.create"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkOrderType Setter
// 工单类型
func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetWorkOrderType(workOrderType string) error {
    r.workOrderType = workOrderType
    r.Set("work_order_type", workOrderType)
    return nil
}

// WorkOrderType Getter
func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetWorkOrderType() string {
    return r.workOrderType
}
// BizType Setter
// 业务类型
func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetBizType() string {
    return r.bizType
}
// Memo Setter
// 工单创建备注
func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

// Memo Getter
func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetMemo() string {
    return r.memo
}
// MemberId Setter
// 货主商家用户id
func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetMemberId(memberId string) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

// MemberId Getter
func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetMemberId() string {
    return r.memberId
}
// MemberRole Setter
// 货主用户角色
func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetMemberRole(memberRole string) error {
    r.memberRole = memberRole
    r.Set("member_role", memberRole)
    return nil
}

// MemberRole Getter
func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetMemberRole() string {
    return r.memberRole
}
// CreatorId Setter
// 创建者淘宝id（区分子账号）
func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetCreatorId(creatorId string) error {
    r.creatorId = creatorId
    r.Set("creator_id", creatorId)
    return nil
}

// CreatorId Getter
func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetCreatorId() string {
    return r.creatorId
}
// CreatorRole Setter
// 创建者角色
func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetCreatorRole(creatorRole string) error {
    r.creatorRole = creatorRole
    r.Set("creator_role", creatorRole)
    return nil
}

// CreatorRole Getter
func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetCreatorRole() string {
    return r.creatorRole
}
// BizEntityValue Setter
// 外部业务系统主键
func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetBizEntityValue(bizEntityValue string) error {
    r.bizEntityValue = bizEntityValue
    r.Set("biz_entity_value", bizEntityValue)
    return nil
}

// BizEntityValue Getter
func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetBizEntityValue() string {
    return r.bizEntityValue
}
// ShopUserId Setter
// 店铺用户id
func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetShopUserId(shopUserId string) error {
    r.shopUserId = shopUserId
    r.Set("shop_user_id", shopUserId)
    return nil
}

// ShopUserId Getter
func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetShopUserId() string {
    return r.shopUserId
}
// TradeId Setter
// 交易订单id
func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetTradeId(tradeId string) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

// TradeId Getter
func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetTradeId() string {
    return r.tradeId
}
// Source Setter
// 工单来源
func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetSource() string {
    return r.source
}
// SourceSign Setter
// 来源签名，用于唯一区分不同的来源方
func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetSourceSign(sourceSign string) error {
    r.sourceSign = sourceSign
    r.Set("source_sign", sourceSign)
    return nil
}

// SourceSign Getter
func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetSourceSign() string {
    return r.sourceSign
}
// MailNo Setter
// 运单号
func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetMailNo(mailNo string) error {
    r.mailNo = mailNo
    r.Set("mail_no", mailNo)
    return nil
}

// MailNo Getter
func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetMailNo() string {
    return r.mailNo
}
// Features Setter
// 扩展字段
func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetFeatures(features string) error {
    r.features = features
    r.Set("features", features)
    return nil
}

// Features Getter
func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetFeatures() string {
    return r.features
}
// AttachPathList Setter
// 凭证地址列表
func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetAttachPathList(attachPathList []string) error {
    r.attachPathList = attachPathList
    r.Set("attach_path_list", attachPathList)
    return nil
}

// AttachPathList Getter
func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetAttachPathList() []string {
    return r.attachPathList
}
