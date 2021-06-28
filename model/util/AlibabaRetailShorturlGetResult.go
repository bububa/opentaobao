package util

// AlibabaRetailShorturlGetResult 
type AlibabaRetailShorturlGetResult struct {

    // module
    
    Module   *ShortUrlDto `json:"module,omitempty" xml:"module,omitempty"`
    

    // errorInfos
    
    ErrorInfos   []ErrorInfo `json:"error_infos,omitempty" xml:"error_infos,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
