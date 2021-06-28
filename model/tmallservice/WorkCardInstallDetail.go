package tmallservice

// WorkCardInstallDetail 
/* model for simplify = false
type WorkCardInstallDetail struct {

    // 机器条码，如果有多个机器码，用英文逗号","隔开
    
    Sn   string `json:"sn,omitempty"`
    

    // 机器安装状态(1未完成；20暂不安装；5完成；12取消)
    
    Status   int64 `json:"status,omitempty"`
    

    // 安装图片，多个图片链接用英文逗号","隔开
    
    ImgUrls   string `json:"img_urls,omitempty"`
    

    // 备注
    
    Memo   string `json:"memo,omitempty"`
    

    // 辅材项目
    
    AccessoryInfo   string `json:"accessory_info,omitempty"`
    

    // 收费金额
    
    InstallFee   string `json:"install_fee,omitempty"`
    

}
*/

// WorkCardInstallDetail 
type WorkCardInstallDetail struct {

    // 机器条码，如果有多个机器码，用英文逗号","隔开
    Sn   string `json:"sn,omitempty"`

    // 机器安装状态(1未完成；20暂不安装；5完成；12取消)
    Status   int64 `json:"status,omitempty"`

    // 安装图片，多个图片链接用英文逗号","隔开
    ImgUrls   string `json:"img_urls,omitempty"`

    // 备注
    Memo   string `json:"memo,omitempty"`

    // 辅材项目
    AccessoryInfo   string `json:"accessory_info,omitempty"`

    // 收费金额
    InstallFee   string `json:"install_fee,omitempty"`

}
