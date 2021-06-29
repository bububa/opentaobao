package alihealthmdeer

// SynArticleInfo 
type SynArticleInfo struct {

    // 作者电话
    
    PhoneNumber   string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
    

    // 作者简介
    
    AuthorIntroduction   string `json:"author_introduction,omitempty" xml:"author_introduction,omitempty"`
    

    // 作者科室
    
    AuthorDepartment   string `json:"author_department,omitempty" xml:"author_department,omitempty"`
    

    // 作者级别
    
    AuthorLevel   string `json:"author_level,omitempty" xml:"author_level,omitempty"`
    

    // 医院级别
    
    HospitalLevel   string `json:"hospital_level,omitempty" xml:"hospital_level,omitempty"`
    

    // 医院名称
    
    HospitalName   string `json:"hospital_name,omitempty" xml:"hospital_name,omitempty"`
    

    // 作者头图
    
    PortraitUrl   string `json:"portrait_url,omitempty" xml:"portrait_url,omitempty"`
    

    // 文章作者名称
    
    AuthorName   string `json:"author_name,omitempty" xml:"author_name,omitempty"`
    

    // 文章正文
    
    Content   string `json:"content,omitempty" xml:"content,omitempty"`
    

    // 原文链接
    
    OriginalArticleUrl   string `json:"original_article_url,omitempty" xml:"original_article_url,omitempty"`
    

    // 文章头图
    
    TitleImage   string `json:"title_image,omitempty" xml:"title_image,omitempty"`
    

    // 文章摘要
    
    ArticleAbstract   string `json:"article_abstract,omitempty" xml:"article_abstract,omitempty"`
    

    // 文章标题
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 文章ID
    
    ArticleId   int64 `json:"article_id,omitempty" xml:"article_id,omitempty"`
    

}
