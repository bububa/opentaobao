package tmallgenie

// DeviceCorpusTopDTO 
type DeviceCorpusTopDTO struct {
    // 操作语料
    CorpusList   []string `json:"corpus_list,omitempty" xml:"corpus_list>string,omitempty"`
    // 支持的操作类型
    FunctionName   string `json:"function_name,omitempty" xml:"function_name,omitempty"`
}
