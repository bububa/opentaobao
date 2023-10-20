package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugUploadExtinfoResultModel 结构体
type AlibabaAlihealthDrugUploadExtinfoResultModel struct {
	// 编码值
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 中文解释
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 接口响应是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务处理是否成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

var poolAlibabaAlihealthDrugUploadExtinfoResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugUploadExtinfoResultModel)
	},
}

// GetAlibabaAlihealthDrugUploadExtinfoResultModel() 从对象池中获取AlibabaAlihealthDrugUploadExtinfoResultModel
func GetAlibabaAlihealthDrugUploadExtinfoResultModel() *AlibabaAlihealthDrugUploadExtinfoResultModel {
	return poolAlibabaAlihealthDrugUploadExtinfoResultModel.Get().(*AlibabaAlihealthDrugUploadExtinfoResultModel)
}

// ReleaseAlibabaAlihealthDrugUploadExtinfoResultModel 释放AlibabaAlihealthDrugUploadExtinfoResultModel
func ReleaseAlibabaAlihealthDrugUploadExtinfoResultModel(v *AlibabaAlihealthDrugUploadExtinfoResultModel) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	v.Model = false
	poolAlibabaAlihealthDrugUploadExtinfoResultModel.Put(v)
}
