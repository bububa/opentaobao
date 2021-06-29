package drug

// TopAlihealthDrugInstructionDTO 
type TopAlihealthDrugInstructionDTO struct {
    // 功能主治
    Purpose   string `json:"purpose,omitempty" xml:"purpose,omitempty"`
    // 禁忌
    Taboo   string `json:"taboo,omitempty" xml:"taboo,omitempty"`
    // 注意事项
    Notes   string `json:"notes,omitempty" xml:"notes,omitempty"`
    // 儿童提醒
    ChildReminder   string `json:"child_reminder,omitempty" xml:"child_reminder,omitempty"`
    // 孕妇提醒
    GravidaReminder   string `json:"gravida_reminder,omitempty" xml:"gravida_reminder,omitempty"`
    // 老人提醒
    OldManReminder   string `json:"old_man_reminder,omitempty" xml:"old_man_reminder,omitempty"`
    // 不良反应
    Adr   string `json:"adr,omitempty" xml:"adr,omitempty"`
    // 药物成分
    Ingredient   string `json:"ingredient,omitempty" xml:"ingredient,omitempty"`
    // 药品性状
    Traits   string `json:"traits,omitempty" xml:"traits,omitempty"`
    // 药物相互作用
    Interaction   string `json:"interaction,omitempty" xml:"interaction,omitempty"`
    // 药理作用
    PharmacologicalEffects   string `json:"pharmacological_effects,omitempty" xml:"pharmacological_effects,omitempty"`
    // 用法用量
    Dosage   string `json:"dosage,omitempty" xml:"dosage,omitempty"`
}
