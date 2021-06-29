package nrt

// MemberSynResponse 
type MemberSynResponse struct {
    // 加密后的taoId
    OpenId   string `json:"open_id,omitempty" xml:"open_id,omitempty"`
    // 业务code
    BizCode   string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
    // 外部会员ID
    OutMemberId   string `json:"out_member_id,omitempty" xml:"out_member_id,omitempty"`
    // sellerid
    SellerId   int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}
