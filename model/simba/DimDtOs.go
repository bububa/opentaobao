package simba

// DimDtOs 
/* model for simplify = false
type DimDtOs struct {

    // tagList
    
    TagList  struct {
        TagOptions  []TagOptions `json:"tag_options,omitempty"`
    } `json:"tag_list,omitempty"`
    

    // 维度名称,如性别,年龄
    
    Name   string `json:"name,omitempty"`
    

    // 维度id,如性别年龄的id
    
    Id   int64 `json:"id,omitempty"`
    

}
*/

// DimDtOs 
type DimDtOs struct {

    // tagList
    TagList   []TagOptions `json:"tag_list,omitempty"`

    // 维度名称,如性别,年龄
    Name   string `json:"name,omitempty"`

    // 维度id,如性别年龄的id
    Id   int64 `json:"id,omitempty"`

}
