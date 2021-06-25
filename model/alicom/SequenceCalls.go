package alicom

// SequenceCalls 
type SequenceCalls struct {

    // 主叫放音
    CallNoPlayCode   string `json:"call_no_play_code,omitempty"`

    // 被叫号码
    CalledNo   string `json:"called_no,omitempty"`

    // 被叫号显
    CalledDisplayNo   string `json:"called_display_no,omitempty"`

    // 被叫放音
    CalledNoPlayCode   string `json:"called_no_play_code,omitempty"`

}
