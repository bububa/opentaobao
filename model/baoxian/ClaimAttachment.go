package baoxian

// ClaimAttachment 
/* model for simplify = false
type ClaimAttachment struct {

    // 附件类型
    
    Type   int64 `json:"type,omitempty"`
    

    // 文件名称
    
    Name   string `json:"name,omitempty"`
    

    // 文件大小
    
    Size   int64 `json:"size,omitempty"`
    

    // 文件路径
    
    Path   string `json:"path,omitempty"`
    

    // 文件描述
    
    Description   string `json:"description,omitempty"`
    

    // 文件类型
    
    FileType   string `json:"file_type,omitempty"`
    

}
*/

// ClaimAttachment 
type ClaimAttachment struct {

    // 附件类型
    Type   int64 `json:"type,omitempty"`

    // 文件名称
    Name   string `json:"name,omitempty"`

    // 文件大小
    Size   int64 `json:"size,omitempty"`

    // 文件路径
    Path   string `json:"path,omitempty"`

    // 文件描述
    Description   string `json:"description,omitempty"`

    // 文件类型
    FileType   string `json:"file_type,omitempty"`

}
