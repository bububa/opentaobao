package tbk

// TaobaoTbkScRelationRecordMapData 结构体
type TaobaoTbkScRelationRecordMapData struct {
	// 带授权的备案链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 已废弃，请调用淘口令生成接口
	PasswordSimple string `json:"password_simple,omitempty" xml:"password_simple,omitempty"`
	// 已废弃，请调用淘口令生成接口
	Model string `json:"model,omitempty" xml:"model,omitempty"`
}
