package alihealth2

// IMRule 结构体
type IMRule struct {
	// INTEGRITY(&#34;对话完整性&#34;),     ATTITUDE(&#34;态度&#34;),     CONTENT_SECURITY(&#34;内容安全&#34;)
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 只有ATTITUDE类型才有值，识别的结果，如：&#34;语气问题&#34;, &#34;辱骂骚扰&#34;，&#34;自我介绍&#34;, &#34;感谢用语&#34;, &#34;祝福用语&#34;,      * &#34;礼貌用语&#34;, &#34;患者安抚&#34;, &#34;随访追问&#34;, &#34;邀评关注&#34;, &#34;平台推荐&#34;等
	Label string `json:"label,omitempty" xml:"label,omitempty"`
	// 1: 违规， 2: 疑似， -1: 该类型下，不对结果进行评判
	Res int64 `json:"res,omitempty" xml:"res,omitempty"`
}
