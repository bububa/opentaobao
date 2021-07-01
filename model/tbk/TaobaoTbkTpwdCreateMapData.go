package tbk

// TaobaoTbkTpwdCreateMapData 结构体
type TaobaoTbkTpwdCreateMapData struct {
	// 非苹果ios14以上版本的设备（即其他ios版本、Android系统等），可以用此淘口令正常在复制到手淘打开
	PasswordSimple string `json:"password_simple,omitempty" xml:"password_simple,omitempty"`
	// 针对苹果ios14及以上版本的苹果设备，手淘将按照示例值信息格式读取淘口令(需包含：数字+羊角符+url，识别规则可能根据ios情况变更)。如需更改淘口令内文案、url等内容，请务必先验证更改后的淘口令在手淘可被识别打开！
	Model string `json:"model,omitempty" xml:"model,omitempty"`
}
