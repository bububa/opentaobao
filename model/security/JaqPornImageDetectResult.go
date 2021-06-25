package security

// JaqPornImageDetectResult 
type JaqPornImageDetectResult struct {

    // 0~100范围的一个浮点值，越接近100，表示色情图像的概率越高（精度到小数点后2位）
    Label   int64 `json:"label,omitempty"`

    // 0表示正常，1表示色情，2表示未确认。用户可以根据自己的场景采信这个值（注：根据图片样本的不断积累，动态调整建议值的阈值设定）
    Rate   string `json:"rate,omitempty"`

}
