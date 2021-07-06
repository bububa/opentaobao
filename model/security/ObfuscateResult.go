package security

// ObfuscateResult 结构体
type ObfuscateResult struct {
	// 总混淆率
	ObfuscatedPercent string `json:"obfuscated_percent,omitempty" xml:"obfuscated_percent,omitempty"`
	// 总的类数量
	TotalClasses int64 `json:"total_classes,omitempty" xml:"total_classes,omitempty"`
	// 总的成员变量数量
	TotalFields int64 `json:"total_fields,omitempty" xml:"total_fields,omitempty"`
	// 总的方法数量
	TotalMethods int64 `json:"total_methods,omitempty" xml:"total_methods,omitempty"`
	// 混淆类的数量
	ObfuscatedClasses int64 `json:"obfuscated_classes,omitempty" xml:"obfuscated_classes,omitempty"`
	// 混淆成员变量的数量
	ObfuscatedFields int64 `json:"obfuscated_fields,omitempty" xml:"obfuscated_fields,omitempty"`
	// 混淆方法的数量
	ObfuscatedMethods int64 `json:"obfuscated_methods,omitempty" xml:"obfuscated_methods,omitempty"`
}
