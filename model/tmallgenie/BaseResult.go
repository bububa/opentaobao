package tmallgenie

// BaseResult 
type BaseResult struct {
    // 实体更新结果对象
    RetValue   *SimpleTextImportResult `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
    // 服务返回错误码
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
    // 服务返回错误信息
    RetMsg   string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
    // [     {       "skillId": 123,       "invocationName": "来个鸟叫",       "name": "t2",       "serviceProviders": null,       "botId": 10,       "class": "com.alibaba.ai.platform.biz.domain.BotSkillsRelInfo",       "iconImgUrl": "//arplatform.alicdn.com/images/90/1499945738188.png",       "longDesc": "2劳动节粉丝撒 uv 那 v 那女 i 啊恶女怕任何 v 去 u 却认为起恢复健康IE肌肤 i 啊viu 话题 uv青海湖去任何欺骗 v额往日 u 问啊好热v 好"     }   ]
    RetValues   *BotSkillsRelInfo `json:"ret_values,omitempty" xml:"ret_values,omitempty"`
}
