package drug

import (
	"sync"
)

// TopAlihealthDrugInstructionDto 结构体
type TopAlihealthDrugInstructionDto struct {
	// 功能主治
	Purpose string `json:"purpose,omitempty" xml:"purpose,omitempty"`
	// 禁忌
	Taboo string `json:"taboo,omitempty" xml:"taboo,omitempty"`
	// 注意事项
	Notes string `json:"notes,omitempty" xml:"notes,omitempty"`
	// 儿童提醒
	ChildReminder string `json:"child_reminder,omitempty" xml:"child_reminder,omitempty"`
	// 孕妇提醒
	GravidaReminder string `json:"gravida_reminder,omitempty" xml:"gravida_reminder,omitempty"`
	// 老人提醒
	OldManReminder string `json:"old_man_reminder,omitempty" xml:"old_man_reminder,omitempty"`
	// 不良反应
	Adr string `json:"adr,omitempty" xml:"adr,omitempty"`
	// 药物成分
	Ingredient string `json:"ingredient,omitempty" xml:"ingredient,omitempty"`
	// 药品性状
	Traits string `json:"traits,omitempty" xml:"traits,omitempty"`
	// 药物相互作用
	Interaction string `json:"interaction,omitempty" xml:"interaction,omitempty"`
	// 药理作用
	PharmacologicalEffects string `json:"pharmacological_effects,omitempty" xml:"pharmacological_effects,omitempty"`
	// 用法用量
	Dosage string `json:"dosage,omitempty" xml:"dosage,omitempty"`
}

var poolTopAlihealthDrugInstructionDto = sync.Pool{
	New: func() any {
		return new(TopAlihealthDrugInstructionDto)
	},
}

// GetTopAlihealthDrugInstructionDto() 从对象池中获取TopAlihealthDrugInstructionDto
func GetTopAlihealthDrugInstructionDto() *TopAlihealthDrugInstructionDto {
	return poolTopAlihealthDrugInstructionDto.Get().(*TopAlihealthDrugInstructionDto)
}

// ReleaseTopAlihealthDrugInstructionDto 释放TopAlihealthDrugInstructionDto
func ReleaseTopAlihealthDrugInstructionDto(v *TopAlihealthDrugInstructionDto) {
	v.Purpose = ""
	v.Taboo = ""
	v.Notes = ""
	v.ChildReminder = ""
	v.GravidaReminder = ""
	v.OldManReminder = ""
	v.Adr = ""
	v.Ingredient = ""
	v.Traits = ""
	v.Interaction = ""
	v.PharmacologicalEffects = ""
	v.Dosage = ""
	poolTopAlihealthDrugInstructionDto.Put(v)
}
