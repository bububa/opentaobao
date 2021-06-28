package mtop

// UploadTokenRequestV 
type UploadTokenRequestV struct {

    // 文件大小，单位byte
    
    FileSize   int64 `json:"file_size,omitempty" xml:"file_size,omitempty"`
    

    // 自定义数据
    
    PrivateData   string `json:"private_data,omitempty" xml:"private_data,omitempty"`
    

    // 上传类型：resumable 或 normal
    
    UploadType   string `json:"upload_type,omitempty" xml:"upload_type,omitempty"`
    

    // 文件内容的CRC32校验和
    
    Crc   int64 `json:"crc,omitempty" xml:"crc,omitempty"`
    

    // 客户端网络类型 wifi 或 2g 或 3g 或 cdma 或 gprs
    
    ClientNetType   string `json:"client_net_type,omitempty" xml:"client_net_type,omitempty"`
    

    // 文件名
    
    FileName   string `json:"file_name,omitempty" xml:"file_name,omitempty"`
    

    // 多媒体中心分配的业务码, mtopupload或其他
    
    BizCode   string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
    

}
