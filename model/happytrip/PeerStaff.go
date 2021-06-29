package happytrip

// PeerStaff 
type PeerStaff struct {
    // 同行人用户id，阿里工号
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 同行人姓名
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}
