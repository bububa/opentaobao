package waybill

// WaybillCloudPrintUpdateRequest 
type WaybillCloudPrintUpdateRequest struct {

    // 物流公司CODE
    
    CpCode   string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
    

    // 物流服务内容<a href="http://open.taobao.com/doc2/detail.htm?spm=a219a.7629140.0.0.eK8aZm&treeId=17&articleId=26765&docType=2">链接</a>
    
    LogisticsServices   string `json:"logistics_services,omitempty" xml:"logistics_services,omitempty"`
    

    // 包裹信息
    
    PackageInfo   *PackageInfoDto `json:"package_info,omitempty" xml:"package_info,omitempty"`
    

    // 收件信息
    
    Recipient   *UserInfoDto `json:"recipient,omitempty" xml:"recipient,omitempty"`
    

    // 发件信息
    
    Sender   *UserInfoDto `json:"sender,omitempty" xml:"sender,omitempty"`
    

    // 模板URL
    
    TemplateUrl   string `json:"template_url,omitempty" xml:"template_url,omitempty"`
    

    // 面单号
    
    WaybillCode   string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
    

    // 请求表示id
    
    ObjectId   string `json:"object_id,omitempty" xml:"object_id,omitempty"`
    

}
