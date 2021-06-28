package security

// OfficialAppVerifyResponse 
/* model for simplify = false
type OfficialAppVerifyResponse struct {

    // 查询任务状态0-处理中 1-处理完成
    
    Status   int64 `json:"status,omitempty"`
    

    // 任务轮询间隔,单位毫秒
    
    QueryInterval   int64 `json:"query_interval,omitempty"`
    

    // 0-未知 1-官方正式 2-官方开发 3-非官方应用 4-待定
    
    OfficialResult   int64 `json:"official_result,omitempty"`
    

    // 包名
    
    PkgName   string `json:"pkg_name,omitempty"`
    

    // 应用名
    
    AppName   string `json:"app_name,omitempty"`
    

    // 证书md5
    
    CertMd5   string `json:"cert_md5,omitempty"`
    

    // 开发者
    
    Developer   string `json:"developer,omitempty"`
    

    // message
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 错误码 400-参数错误  500-服务错误
    
    ErrCode   int64 `json:"err_code,omitempty"`
    

    // 请求是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// OfficialAppVerifyResponse 
type OfficialAppVerifyResponse struct {

    // 查询任务状态0-处理中 1-处理完成
    Status   int64 `json:"status,omitempty"`

    // 任务轮询间隔,单位毫秒
    QueryInterval   int64 `json:"query_interval,omitempty"`

    // 0-未知 1-官方正式 2-官方开发 3-非官方应用 4-待定
    OfficialResult   int64 `json:"official_result,omitempty"`

    // 包名
    PkgName   string `json:"pkg_name,omitempty"`

    // 应用名
    AppName   string `json:"app_name,omitempty"`

    // 证书md5
    CertMd5   string `json:"cert_md5,omitempty"`

    // 开发者
    Developer   string `json:"developer,omitempty"`

    // message
    ErrMsg   string `json:"err_msg,omitempty"`

    // 错误码 400-参数错误  500-服务错误
    ErrCode   int64 `json:"err_code,omitempty"`

    // 请求是否成功
    Success   bool `json:"success,omitempty"`

}
