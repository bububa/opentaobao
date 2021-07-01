package drugtrace

// AlibabaAlihealthDrugCodeCommonListCodeinfoResultModel 
type AlibabaAlihealthDrugCodeCommonListCodeinfoResultModel struct {
    // 内层大对象
    Models   []CodeFullInfoDto `json:"models,omitempty" xml:"models>code_full_info_dto,omitempty"`
    // 消息码
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 消息提示内容
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 查询成功失败标记
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
