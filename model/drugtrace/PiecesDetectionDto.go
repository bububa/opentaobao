package drugtrace

import (
	"sync"
)

// PiecesDetectionDto 结构体
type PiecesDetectionDto struct {
	// 饮片检验报告书（上传图片）图片建议尺寸：height: 310px;width: 670px;
	InspectionReportPictures []string `json:"inspection_report_pictures,omitempty" xml:"inspection_report_pictures>string,omitempty"`
	// 饮片图片（上传图片）图片建议尺寸：height: 310px;width: 670px;
	PiecesPictures []string `json:"pieces_pictures,omitempty" xml:"pieces_pictures>string,omitempty"`
	// 农药残留检测
	PesticidesDetection string `json:"pesticides_detection,omitempty" xml:"pesticides_detection,omitempty"`
	// 重金属及有害元素检测
	PiecesHarmDetection string `json:"pieces_harm_detection,omitempty" xml:"pieces_harm_detection,omitempty"`
	// 黄曲霉毒素检测
	AflatoxinDetection string `json:"aflatoxin_detection,omitempty" xml:"aflatoxin_detection,omitempty"`
	// 饮片执行标准
	ExecStandard string `json:"exec_standard,omitempty" xml:"exec_standard,omitempty"`
}

var poolPiecesDetectionDto = sync.Pool{
	New: func() any {
		return new(PiecesDetectionDto)
	},
}

// GetPiecesDetectionDto() 从对象池中获取PiecesDetectionDto
func GetPiecesDetectionDto() *PiecesDetectionDto {
	return poolPiecesDetectionDto.Get().(*PiecesDetectionDto)
}

// ReleasePiecesDetectionDto 释放PiecesDetectionDto
func ReleasePiecesDetectionDto(v *PiecesDetectionDto) {
	v.InspectionReportPictures = v.InspectionReportPictures[:0]
	v.PiecesPictures = v.PiecesPictures[:0]
	v.PesticidesDetection = ""
	v.PiecesHarmDetection = ""
	v.AflatoxinDetection = ""
	v.ExecStandard = ""
	poolPiecesDetectionDto.Put(v)
}
