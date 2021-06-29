package shenjing

// PostObjectPolicyVo 
type PostObjectPolicyVo struct {
    // 阿里云RAM账号的accessId
    Accessid   string `json:"accessid,omitempty" xml:"accessid,omitempty"`
    // 加密过的上传策略
    Policy   string `json:"policy,omitempty" xml:"policy,omitempty"`
    // 上传策略签名
    Signature   string `json:"signature,omitempty" xml:"signature,omitempty"`
    // 文件上传目录
    Dir   string `json:"dir,omitempty" xml:"dir,omitempty"`
    // 上传策略有效时间
    Expire   string `json:"expire,omitempty" xml:"expire,omitempty"`
    // oss的bucket访问路径
    Host   string `json:"host,omitempty" xml:"host,omitempty"`
}
