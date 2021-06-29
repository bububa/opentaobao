package tmallgenie

// SimpleTextImportResult 
type SimpleTextImportResult struct {

    // 已存在或重复的实体数
    
    NumExist   int64 `json:"num_exist,omitempty" xml:"num_exist,omitempty"`
    

    // 上传失败的实体数
    
    NumFailed   int64 `json:"num_failed,omitempty" xml:"num_failed,omitempty"`
    

    // 上传成功的实体数
    
    NumSuccessful   int64 `json:"num_successful,omitempty" xml:"num_successful,omitempty"`
    

}
