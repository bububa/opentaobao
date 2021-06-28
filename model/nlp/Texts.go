package nlp

// Texts 
/* model for simplify = false
type Texts struct {

    // 业务处理ID
    
    Id   string `json:"id,omitempty"`
    

    // 文本相似度匹配文本
    
    ContentSrc   string `json:"content_src,omitempty"`
    

    // 文本相似度匹配文本内容模板
    
    ContentDst   string `json:"content_dst,omitempty"`
    

    // 文本相似度匹配类型：1为余弦距离，2为编辑距离，3为simHash距离
    
    Type   int64 `json:"type,omitempty"`
    

}
*/

// Texts 
type Texts struct {

    // 业务处理ID
    Id   string `json:"id,omitempty"`

    // 文本相似度匹配文本
    ContentSrc   string `json:"content_src,omitempty"`

    // 文本相似度匹配文本内容模板
    ContentDst   string `json:"content_dst,omitempty"`

    // 文本相似度匹配类型：1为余弦距离，2为编辑距离，3为simHash距离
    Type   int64 `json:"type,omitempty"`

}
