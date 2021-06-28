package util

// Checkpoints 
type Checkpoints struct {

    // 检查的场景。antispam为黄暴政
    
    Scene   string `json:"scene,omitempty" xml:"scene,omitempty"`
    

    // 检查的场景标签。normal：正常文本 spam：含垃圾信息；ad：广告；politics：涉政；terrorism：暴恐；abuse：辱骂；porn：色情；flood：灌水；contraband：违禁；meaningless：无意义
    
    Label   string `json:"label,omitempty" xml:"label,omitempty"`
    

    // 结果建议
    
    Suggestion   string `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
    

    // 结果准确度
    
    Rate   string `json:"rate,omitempty" xml:"rate,omitempty"`
    

}
