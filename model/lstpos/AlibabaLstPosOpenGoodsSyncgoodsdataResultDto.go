package lstpos

// AlibabaLstPosOpenGoodsSyncgoodsdataResultDto 
type AlibabaLstPosOpenGoodsSyncgoodsdataResultDto struct {
    // 错误信息描述
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 接口具体返回的业务数据对象
    Modules   []ErrorResult `json:"modules,omitempty" xml:"modules>error_result,omitempty"`
    // 业务错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 接口调用是否成功 true:调用成功 false:调用失败
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
