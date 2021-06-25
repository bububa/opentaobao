package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟工单创建接口 APIRequest
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
    attachPathList   []String 

}

func NewCainiaoCbossWorkplatformWorkorderCreateRequest() *CainiaoCbossWorkplatformWorkorderCreateRequest{
    return &CainiaoCbossWorkplatformWorkorderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetApiMethodName() string {
    return "cainiao.cboss.workplatform.workorder.create"
}

func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetWorkOrderType(workOrderType string) error {
    r.workOrderType = workOrderType
    r.Set("work_order_type", workOrderType)
    return nil
}

func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetWorkOrderType() string {
    return r.workOrderType
}

func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetBizType() string {
    return r.bizType
}

func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetMemo() string {
    return r.memo
}

func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetMemberId(memberId string) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetMemberId() string {
    return r.memberId
}

func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetMemberRole(memberRole string) error {
    r.memberRole = memberRole
    r.Set("member_role", memberRole)
    return nil
}

func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetMemberRole() string {
    return r.memberRole
}

func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetCreatorId(creatorId string) error {
    r.creatorId = creatorId
    r.Set("creator_id", creatorId)
    return nil
}

func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetCreatorId() string {
    return r.creatorId
}

func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetCreatorRole(creatorRole string) error {
    r.creatorRole = creatorRole
    r.Set("creator_role", creatorRole)
    return nil
}

func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetCreatorRole() string {
    return r.creatorRole
}

func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetBizEntityValue(bizEntityValue string) error {
    r.bizEntityValue = bizEntityValue
    r.Set("biz_entity_value", bizEntityValue)
    return nil
}

func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetBizEntityValue() string {
    return r.bizEntityValue
}

func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetShopUserId(shopUserId string) error {
    r.shopUserId = shopUserId
    r.Set("shop_user_id", shopUserId)
    return nil
}

func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetShopUserId() string {
    return r.shopUserId
}

func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetTradeId(tradeId string) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetTradeId() string {
    return r.tradeId
}

func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetSource() string {
    return r.source
}

func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetSourceSign(sourceSign string) error {
    r.sourceSign = sourceSign
    r.Set("source_sign", sourceSign)
    return nil
}

func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetSourceSign() string {
    return r.sourceSign
}

func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetMailNo(mailNo string) error {
    r.mailNo = mailNo
    r.Set("mail_no", mailNo)
    return nil
}

func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetMailNo() string {
    return r.mailNo
}

func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetFeatures(features string) error {
    r.features = features
    r.Set("features", features)
    return nil
}

func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetFeatures() string {
    return r.features
}

func (r *CainiaoCbossWorkplatformWorkorderCreateRequest) SetAttachPathList(attachPathList []String) error {
    r.attachPathList = attachPathList
    r.Set("attach_path_list", attachPathList)
    return nil
}

func (r CainiaoCbossWorkplatformWorkorderCreateRequest) GetAttachPathList() []String {
    return r.attachPathList
}

