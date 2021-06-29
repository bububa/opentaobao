package nrt

// NrtMemberDTO 
type NrtMemberDTO struct {

    // 手机号
    
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    

    // 外部会员ID
    
    OutMemberId   string `json:"out_member_id,omitempty" xml:"out_member_id,omitempty"`
    

    // 同步时间
    
    GmtModified   int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    

    // 业务code
    
    BizCode   string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
    

    // 淘宝ID
    
    OpenId   string `json:"open_id,omitempty" xml:"open_id,omitempty"`
    

    // 操作类型
    
    OpType   string `json:"op_type,omitempty" xml:"op_type,omitempty"`
    

    // 会员等级
    
    LevelName   string `json:"level_name,omitempty" xml:"level_name,omitempty"`
    

    // 版本号
    
    Version   int64 `json:"version,omitempty" xml:"version,omitempty"`
    

    // 会员名称
    
    RealName   string `json:"real_name,omitempty" xml:"real_name,omitempty"`
    

    // 幂等ID
    
    OutId   string `json:"out_id,omitempty" xml:"out_id,omitempty"`
    

    // 会员地址
    
    Addr   string `json:"addr,omitempty" xml:"addr,omitempty"`
    

    // 状态 1:正常 0:冻结
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 淘宝nick
    
    TaoNick   string `json:"tao_nick,omitempty" xml:"tao_nick,omitempty"`
    

}
