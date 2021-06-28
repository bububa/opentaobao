package wlb

// WlbSellerSubscription 
/* model for simplify = false
type WlbSellerSubscription struct {

    // 定购ID
    
    Id   int64 `json:"id,omitempty"`
    

    // 定购用户ID
    
    SubscriberUserId   int64 `json:"subscriber_user_id,omitempty"`
    

    // 定购用户NICK
    
    SubscriberUserNick   string `json:"subscriber_user_nick,omitempty"`
    

    // 服务商ID
    
    ProviderUserId   int64 `json:"provider_user_id,omitempty"`
    

    // 服务ID
    
    ServiceId   int64 `json:"service_id,omitempty"`
    

    // 定购有效开始日期
    
    StartDate   string `json:"start_date,omitempty"`
    

    // 定购有效结束日期
    
    EndDate   string `json:"end_date,omitempty"`
    

    // 服务编码
    
    ServiceCode   string `json:"service_code,omitempty"`
    

    // 服务名
    
    ServiceName   string `json:"service_name,omitempty"`
    

    // 服务类型，<br/>STORE 1-仓储、<br/>TMS 2-TMS、<br/>PACKAGE 3-包装服务<br/>SUPPLIER 4-供货<br/>INSTALL 5-安装<br/>COMPLEX_SERVICE 100-综合服务
    
    ServiceType   string `json:"service_type,omitempty"`
    

    // 父定购ID<br/>可通过该字段来得之服务上下级关系。<br/>例定购了仓储服务，下又有TMS服务。<br/>该字段保存仓储服务ID。
    
    ParentId   int64 `json:"parent_id,omitempty"`
    

    // 状态<br/>AUDITING 1-待审核<br/>CANCEL 2-撤销<br/>CHECKED 3-审核通过<br/>FAILED 4-审核未通过<br/>SYNCHRONIZING 5-同步中
    
    Status   string `json:"status,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty"`
    

    // 联系人地址信息
    
    WlbPartnerAddress  *struct {
        WlbPartnerAddress  *WlbPartnerAddress `json:"wlb_partner_address,omitempty"`
    } `json:"wlb_partner_address,omitempty"`
    

    // 联系人联系详情
    
    WlbPartnerContact  *struct {
        WlbPartnerContact  *WlbPartnerContact `json:"wlb_partner_contact,omitempty"`
    } `json:"wlb_partner_contact,omitempty"`
    

    // 判断该仓库是否是实体仓，还是虚拟仓，null是实体仓，10:代表虚拟仓
    
    IsOwnService   int64 `json:"is_own_service,omitempty"`
    

    // 自有仓的别名，当当前订购记录为自有仓时才会有值
    
    ServiceAlias   string `json:"service_alias,omitempty"`
    

}
*/

// WlbSellerSubscription 
type WlbSellerSubscription struct {

    // 定购ID
    Id   int64 `json:"id,omitempty"`

    // 定购用户ID
    SubscriberUserId   int64 `json:"subscriber_user_id,omitempty"`

    // 定购用户NICK
    SubscriberUserNick   string `json:"subscriber_user_nick,omitempty"`

    // 服务商ID
    ProviderUserId   int64 `json:"provider_user_id,omitempty"`

    // 服务ID
    ServiceId   int64 `json:"service_id,omitempty"`

    // 定购有效开始日期
    StartDate   string `json:"start_date,omitempty"`

    // 定购有效结束日期
    EndDate   string `json:"end_date,omitempty"`

    // 服务编码
    ServiceCode   string `json:"service_code,omitempty"`

    // 服务名
    ServiceName   string `json:"service_name,omitempty"`

    // 服务类型，<br/>STORE 1-仓储、<br/>TMS 2-TMS、<br/>PACKAGE 3-包装服务<br/>SUPPLIER 4-供货<br/>INSTALL 5-安装<br/>COMPLEX_SERVICE 100-综合服务
    ServiceType   string `json:"service_type,omitempty"`

    // 父定购ID<br/>可通过该字段来得之服务上下级关系。<br/>例定购了仓储服务，下又有TMS服务。<br/>该字段保存仓储服务ID。
    ParentId   int64 `json:"parent_id,omitempty"`

    // 状态<br/>AUDITING 1-待审核<br/>CANCEL 2-撤销<br/>CHECKED 3-审核通过<br/>FAILED 4-审核未通过<br/>SYNCHRONIZING 5-同步中
    Status   string `json:"status,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 联系人地址信息
    WlbPartnerAddress   *WlbPartnerAddress `json:"wlb_partner_address,omitempty"`

    // 联系人联系详情
    WlbPartnerContact   *WlbPartnerContact `json:"wlb_partner_contact,omitempty"`

    // 判断该仓库是否是实体仓，还是虚拟仓，null是实体仓，10:代表虚拟仓
    IsOwnService   int64 `json:"is_own_service,omitempty"`

    // 自有仓的别名，当当前订购记录为自有仓时才会有值
    ServiceAlias   string `json:"service_alias,omitempty"`

}
