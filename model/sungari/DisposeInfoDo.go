package sungari

// DisposeInfoDo 
type DisposeInfoDo struct {

    // 处置类型，1:商品处置  2：经营者处置
    
    ApiType   int64 `json:"api_type,omitempty" xml:"api_type,omitempty"`
    

    // 抽检报告编号
    
    CheckNumber   string `json:"check_number,omitempty" xml:"check_number,omitempty"`
    

    // 处置函单位名称
    
    CompanyName   string `json:"company_name,omitempty" xml:"company_name,omitempty"`
    

    // 处置函录入时间
    
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    

    // 处置对象内容
    
    DisposeContent   string `json:"dispose_content,omitempty" xml:"dispose_content,omitempty"`
    

    // 处置对象类型，1:订单号  2：商品ID  3:商家会员账号 4：全网排查关键字
    
    DisposeType   int64 `json:"dispose_type,omitempty" xml:"dispose_type,omitempty"`
    

    // 处置函件文号
    
    DocName   string `json:"doc_name,omitempty" xml:"doc_name,omitempty"`
    

    // 资质证件号（执照）
    
    LicenceNo   string `json:"licence_no,omitempty" xml:"licence_no,omitempty"`
    

    // 处置函联系人
    
    LinkMan   string `json:"link_man,omitempty" xml:"link_man,omitempty"`
    

    // 处置函件Oss中的key
    
    OssKey   string `json:"oss_key,omitempty" xml:"oss_key,omitempty"`
    

    // 处置函联系人电话
    
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    

    // 处置函件接收平台，1:淘宝  2：天猫  3:1688
    
    PlatformType   int64 `json:"platform_type,omitempty" xml:"platform_type,omitempty"`
    

    // 处置原因
    
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    

    // 备注信息
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 原始函单位地址
    
    SourceAddress   string `json:"source_address,omitempty" xml:"source_address,omitempty"`
    

    // 原始函单位名称
    
    SourceCompanyName   string `json:"source_company_name,omitempty" xml:"source_company_name,omitempty"`
    

    // 原始函联系人
    
    SourceLinkMan   string `json:"source_link_man,omitempty" xml:"source_link_man,omitempty"`
    

    // 原始函件Oss中的key
    
    SourceOssKey   string `json:"source_oss_key,omitempty" xml:"source_oss_key,omitempty"`
    

    // 原始函联系人电话
    
    SourcePhone   string `json:"source_phone,omitempty" xml:"source_phone,omitempty"`
    

    // 原始函单位邮编
    
    SourcePostCode   string `json:"source_post_code,omitempty" xml:"source_post_code,omitempty"`
    

    // 原始函接收时间
    
    SourceTime   string `json:"source_time,omitempty" xml:"source_time,omitempty"`
    

    // 函件url链接
    
    OssKeyUrl   string `json:"oss_key_url,omitempty" xml:"oss_key_url,omitempty"`
    

    // 原始函链接
    
    SourceOssKeyUrl   string `json:"source_oss_key_url,omitempty" xml:"source_oss_key_url,omitempty"`
    

}
