package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytYbGetcoderelationModel 结构体
type AlibabaAlihealthDrugKytYbGetcoderelationModel struct {
	// 码关联关系DTO
	CodeRelationDtoList []CodeRelationDto `json:"code_relation_dto_list,omitempty" xml:"code_relation_dto_list>code_relation_dto,omitempty"`
	// 包装比例
	PkgRatio string `json:"pkg_ratio,omitempty" xml:"pkg_ratio,omitempty"`
}

var poolAlibabaAlihealthDrugKytYbGetcoderelationModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytYbGetcoderelationModel)
	},
}

// GetAlibabaAlihealthDrugKytYbGetcoderelationModel() 从对象池中获取AlibabaAlihealthDrugKytYbGetcoderelationModel
func GetAlibabaAlihealthDrugKytYbGetcoderelationModel() *AlibabaAlihealthDrugKytYbGetcoderelationModel {
	return poolAlibabaAlihealthDrugKytYbGetcoderelationModel.Get().(*AlibabaAlihealthDrugKytYbGetcoderelationModel)
}

// ReleaseAlibabaAlihealthDrugKytYbGetcoderelationModel 释放AlibabaAlihealthDrugKytYbGetcoderelationModel
func ReleaseAlibabaAlihealthDrugKytYbGetcoderelationModel(v *AlibabaAlihealthDrugKytYbGetcoderelationModel) {
	v.CodeRelationDtoList = v.CodeRelationDtoList[:0]
	v.PkgRatio = ""
	poolAlibabaAlihealthDrugKytYbGetcoderelationModel.Put(v)
}
