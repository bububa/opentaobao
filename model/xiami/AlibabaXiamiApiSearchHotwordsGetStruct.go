package xiami

// AlibabaXiamiApiSearchHotwordsGetStruct 
type AlibabaXiamiApiSearchHotwordsGetStruct struct {

    // 星标词
    
    StarWords   []AlibabaXiamiApiSearchHotwordsGetStruct `json:"star_words,omitempty" xml:"star_words,omitempty"`
    

    // 搜索热词
    
    SearchWords   []AlibabaXiamiApiSearchHotwordsGetStruct `json:"search_words,omitempty" xml:"search_words,omitempty"`
    

    // 星标词
    
    Word   string `json:"word,omitempty" xml:"word,omitempty"`
    

    // 跳转url
    
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
    

    // 排位变化
    
    Change   int64 `json:"change,omitempty" xml:"change,omitempty"`
    

}
