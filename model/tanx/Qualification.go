package tanx

// Qualification 
/* model for simplify = false
type Qualification struct {

    // 上传的客户资质元素id
    
    ElementId   int64 `json:"element_id,omitempty"`
    

    // 和资质平台交互的userId
    
    UserId   int64 `json:"user_id,omitempty"`
    

    // 资质失效时间
    
    EndTime   string `json:"end_time,omitempty"`
    

    // 资质生效时间
    
    StartTime   string `json:"start_time,omitempty"`
    

    // 为本次上传客户资质的操作取一个名称,如果为空则系统自动生成一个(选填)
    
    Name   string `json:"name,omitempty"`
    

    // 用户附加内容(明星资质的备注)
    
    Supplement   string `json:"supplement,omitempty"`
    

    // 资质内容,如果是图片请先调用uploadQualificationPic转换成url上传，在填入返回url
    
    UrlContents  struct {
        String  []string `json:"string,omitempty"`
    } `json:"url_contents,omitempty"`
    

    // 本次上传的资质内容id,只有在modify时才需要填写
    
    Id   int64 `json:"id,omitempty"`
    

}
*/

// Qualification 
type Qualification struct {

    // 上传的客户资质元素id
    ElementId   int64 `json:"element_id,omitempty"`

    // 和资质平台交互的userId
    UserId   int64 `json:"user_id,omitempty"`

    // 资质失效时间
    EndTime   string `json:"end_time,omitempty"`

    // 资质生效时间
    StartTime   string `json:"start_time,omitempty"`

    // 为本次上传客户资质的操作取一个名称,如果为空则系统自动生成一个(选填)
    Name   string `json:"name,omitempty"`

    // 用户附加内容(明星资质的备注)
    Supplement   string `json:"supplement,omitempty"`

    // 资质内容,如果是图片请先调用uploadQualificationPic转换成url上传，在填入返回url
    UrlContents   []string `json:"url_contents,omitempty"`

    // 本次上传的资质内容id,只有在modify时才需要填写
    Id   int64 `json:"id,omitempty"`

}
