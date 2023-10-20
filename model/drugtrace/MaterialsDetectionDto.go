package drugtrace

import (
	"sync"
)

// MaterialsDetectionDto 结构体
type MaterialsDetectionDto struct {
	// 药材检验报告书（上传图片）图片建议尺寸：height: 310px;width: 670px;
	InspectionReportPictures []string `json:"inspection_report_pictures,omitempty" xml:"inspection_report_pictures>string,omitempty"`
	// 药材图片（上传图片）图片建议尺寸：height: 310px;width: 670px;
	MaterialsPictures []string `json:"materials_pictures,omitempty" xml:"materials_pictures>string,omitempty"`
	// 农药残留检测
	PesticidesDetection string `json:"pesticides_detection,omitempty" xml:"pesticides_detection,omitempty"`
	// 重金属及有害元素检测
	PiecesHarmDetection string `json:"pieces_harm_detection,omitempty" xml:"pieces_harm_detection,omitempty"`
	// 黄曲霉毒素检测
	AflatoxinDetection string `json:"aflatoxin_detection,omitempty" xml:"aflatoxin_detection,omitempty"`
	// 药材执行标准
	ExecStandard string `json:"exec_standard,omitempty" xml:"exec_standard,omitempty"`
	// 药材批号
	MaterialsBatchNo string `json:"materials_batch_no,omitempty" xml:"materials_batch_no,omitempty"`
}

var poolMaterialsDetectionDto = sync.Pool{
	New: func() any {
		return new(MaterialsDetectionDto)
	},
}

// GetMaterialsDetectionDto() 从对象池中获取MaterialsDetectionDto
func GetMaterialsDetectionDto() *MaterialsDetectionDto {
	return poolMaterialsDetectionDto.Get().(*MaterialsDetectionDto)
}

// ReleaseMaterialsDetectionDto 释放MaterialsDetectionDto
func ReleaseMaterialsDetectionDto(v *MaterialsDetectionDto) {
	v.InspectionReportPictures = v.InspectionReportPictures[:0]
	v.MaterialsPictures = v.MaterialsPictures[:0]
	v.PesticidesDetection = ""
	v.PiecesHarmDetection = ""
	v.AflatoxinDetection = ""
	v.ExecStandard = ""
	v.MaterialsBatchNo = ""
	poolMaterialsDetectionDto.Put(v)
}
