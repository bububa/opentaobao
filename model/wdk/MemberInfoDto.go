package wdk

// MemberInfoDto 
type MemberInfoDto struct {

    // 淘宝用户昵称
    Nick   string `json:"nick,omitempty"`

    // 对应淘宝账号的OpenUID
    UicId   string `json:"uic_id,omitempty"`

}
