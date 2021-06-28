package security

// ScanAppInfo 
/* model for simplify = false
type ScanAppInfo struct {

    // APP数据类型 1-App URL 2-App MD5，应用加固仅支持1，风险扫描支持1和2
    
    DataType   int64 `json:"data_type,omitempty"`
    

    // APP数据  dataType=1时填写 App包的下载地址;  dataType=2时填写 App包的md5值
    
    Data   string `json:"data,omitempty"`
    

    // APP包的MD5值,dataType=1时必填,用于文件校验
    
    Md5   string `json:"md5,omitempty"`
    

    // APP包大小(单位:字节),dataType=1时必填,用于文件校验
    
    Size   int64 `json:"size,omitempty"`
    

    // 任务处理完成后的反向通知回调地址,请不要使用ip地址,可能会无法回调,dataType=1时必填,通知为GET请求,请求URL: callbackUrl+"?item_id=xxx&task_status=1"; item_id为应用加固/风险扫描接口返回的任务ID; task_status为任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时; 对于应用加固,接收到通知后如果task_status为1则可通过对应的查询接口查询加固/扫描结果; 对于应用风险扫描,如果task_status为1,3,4均可通过对应查询接口查询到结果,但不包括失败的扫描项的结果
    
    CallbackUrl   string `json:"callback_url,omitempty"`
    

    // app 类型，1-apk, 2-ipa（暂不支持）
    
    AppOsType   int64 `json:"app_os_type,omitempty"`
    

}
*/

// ScanAppInfo 
type ScanAppInfo struct {

    // APP数据类型 1-App URL 2-App MD5，应用加固仅支持1，风险扫描支持1和2
    DataType   int64 `json:"data_type,omitempty"`

    // APP数据  dataType=1时填写 App包的下载地址;  dataType=2时填写 App包的md5值
    Data   string `json:"data,omitempty"`

    // APP包的MD5值,dataType=1时必填,用于文件校验
    Md5   string `json:"md5,omitempty"`

    // APP包大小(单位:字节),dataType=1时必填,用于文件校验
    Size   int64 `json:"size,omitempty"`

    // 任务处理完成后的反向通知回调地址,请不要使用ip地址,可能会无法回调,dataType=1时必填,通知为GET请求,请求URL: callbackUrl+"?item_id=xxx&task_status=1"; item_id为应用加固/风险扫描接口返回的任务ID; task_status为任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时; 对于应用加固,接收到通知后如果task_status为1则可通过对应的查询接口查询加固/扫描结果; 对于应用风险扫描,如果task_status为1,3,4均可通过对应查询接口查询到结果,但不包括失败的扫描项的结果
    CallbackUrl   string `json:"callback_url,omitempty"`

    // app 类型，1-apk, 2-ipa（暂不支持）
    AppOsType   int64 `json:"app_os_type,omitempty"`

}
