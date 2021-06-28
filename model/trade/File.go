package trade

// File 
/* model for simplify = false
type File struct {

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 返回的是绝对路径如：http://img07.taobaocdn.com/imgextra/i7/22670458/T2dD0kXb4cXXXXXXXX_!!22670458.jpg
    
    FilePath   string `json:"file_path,omitempty"`
    

    // 文件的大小
    
    Size   int64 `json:"size,omitempty"`
    

    // 图片状态,unfroze代表没有被冻结，froze代表被冻结,pass代表排查通过
    
    Status   string `json:"status,omitempty"`
    

    // 文件是否删除
    
    Deleted   string `json:"deleted,omitempty"`
    

    // 图片像素，如果非图片，该值为空
    
    PicturePix   string `json:"picture_pix,omitempty"`
    

}
*/

// File 
type File struct {

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 返回的是绝对路径如：http://img07.taobaocdn.com/imgextra/i7/22670458/T2dD0kXb4cXXXXXXXX_!!22670458.jpg
    FilePath   string `json:"file_path,omitempty"`

    // 文件的大小
    Size   int64 `json:"size,omitempty"`

    // 图片状态,unfroze代表没有被冻结，froze代表被冻结,pass代表排查通过
    Status   string `json:"status,omitempty"`

    // 文件是否删除
    Deleted   string `json:"deleted,omitempty"`

    // 图片像素，如果非图片，该值为空
    PicturePix   string `json:"picture_pix,omitempty"`

}
