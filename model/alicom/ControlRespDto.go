package alicom

// ControlRespDto 
type ControlRespDto struct {

    // 接续控制信息:CONTINUE(接续),REJECT(拦截),IVR(收取用户键盘输入内容)
    ControlOperate   string `json:"control_operate,omitempty"`

    // controlMsg
    ControlMsg   string `json:"control_msg,omitempty"`

    // 对应到小号平台的能力类型:AXB、AXN、AXN_EXTENSION_REUSE(AXN分机复用)
    ProductType   string `json:"product_type,omitempty"`

    // 主叫放音编码
    CallNoPlayCode   string `json:"call_no_play_code,omitempty"`

    // 被叫放音编码
    CalledNoPlayCode   string `json:"called_no_play_code,omitempty"`

    // 对应到阿里侧的绑定信息
    Subs   *Subs `json:"subs,omitempty"`

    // 是否媒体资源降级,放弃录音放音功能；接入方无此相关功能，可忽略
    MediaDegrade   bool `json:"media_degrade,omitempty"`

}
