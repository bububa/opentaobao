package tmallgenie

// BaseResult 结构体
type BaseResult struct {
	// 服务返回错误信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 实体更新结果对象
	RetValue *SimpleTextImportResult `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
	// 服务返回错误码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// [     {       &#34;skillId&#34;: 123,       &#34;invocationName&#34;: &#34;来个鸟叫&#34;,       &#34;name&#34;: &#34;t2&#34;,       &#34;serviceProviders&#34;: null,       &#34;botId&#34;: 10,       &#34;class&#34;: &#34;com.alibaba.ai.platform.biz.domain.BotSkillsRelInfo&#34;,       &#34;iconImgUrl&#34;: &#34;//arplatform.alicdn.com/images/90/1499945738188.png&#34;,       &#34;longDesc&#34;: &#34;2劳动节粉丝撒 uv 那 v 那女 i 啊恶女怕任何 v 去 u 却认为起恢复健康IE肌肤 i 啊viu 话题 uv青海湖去任何欺骗 v额往日 u 问啊好热v 好&#34;     }   ]
	RetValues *BotSkillsRelInfo `json:"ret_values,omitempty" xml:"ret_values,omitempty"`
}
