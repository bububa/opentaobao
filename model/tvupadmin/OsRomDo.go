package tvupadmin

// OsRomDO 
type OsRomDO struct {

    // 主键ID
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 版本
    
    VersionId   int64 `json:"version_id,omitempty" xml:"version_id,omitempty"`
    

    // 基础版本
    
    BaseVersionId   int64 `json:"base_version_id,omitempty" xml:"base_version_id,omitempty"`
    

    // 下载地址
    
    DownloadPath   string `json:"download_path,omitempty" xml:"download_path,omitempty"`
    

    // 下载MD5
    
    Downloadmd5   string `json:"downloadmd5,omitempty" xml:"downloadmd5,omitempty"`
    

    // 大小
    
    Size   string `json:"size,omitempty" xml:"size,omitempty"`
    

    // 分片数
    
    Splitnum   int64 `json:"splitnum,omitempty" xml:"splitnum,omitempty"`
    

    // 删除标识
    
    IsDelete   string `json:"is_delete,omitempty" xml:"is_delete,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModify   string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
    

}
