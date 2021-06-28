package xiami

// AlibabaXiamiApiSearchHotwordsGetStruct 
/* model for simplify = false
type AlibabaXiamiApiSearchHotwordsGetStruct struct {

    // 星标词
    
    StarWords  struct {
        AlibabaXiamiApiSearchHotwordsGetStruct  []AlibabaXiamiApiSearchHotwordsGetStruct `json:"alibaba_xiami_api_search_hotwords_get_struct,omitempty"`
    } `json:"star_words,omitempty"`
    

    // 搜索热词
    
    SearchWords  struct {
        AlibabaXiamiApiSearchHotwordsGetStruct  []AlibabaXiamiApiSearchHotwordsGetStruct `json:"alibaba_xiami_api_search_hotwords_get_struct,omitempty"`
    } `json:"search_words,omitempty"`
    

    // 星标词
    
    Word   string `json:"word,omitempty"`
    

    // 跳转url
    
    Url   string `json:"url,omitempty"`
    

    // 排位变化
    
    Change   int64 `json:"change,omitempty"`
    

}
*/

// AlibabaXiamiApiSearchHotwordsGetStruct 
type AlibabaXiamiApiSearchHotwordsGetStruct struct {

    // 星标词
    StarWords   []AlibabaXiamiApiSearchHotwordsGetStruct `json:"star_words,omitempty"`

    // 搜索热词
    SearchWords   []AlibabaXiamiApiSearchHotwordsGetStruct `json:"search_words,omitempty"`

    // 星标词
    Word   string `json:"word,omitempty"`

    // 跳转url
    Url   string `json:"url,omitempty"`

    // 排位变化
    Change   int64 `json:"change,omitempty"`

}
