package feedflow

// LaunchAreaDto 
type LaunchAreaDto struct {
    // 地域的code
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    // 地域名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
}
