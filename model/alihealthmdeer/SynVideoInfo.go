package alihealthmdeer

// SynVideoInfo 
type SynVideoInfo struct {
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
    // 作者头像
    PortraitUrl   string `json:"portrait_url,omitempty" xml:"portrait_url,omitempty"`
    // 作者名称
    AuthorName   string `json:"author_name,omitempty" xml:"author_name,omitempty"`
    // 视频原链接
    OriginalUrl   string `json:"original_url,omitempty" xml:"original_url,omitempty"`
    // 视频简介
    VideoIntroduction   string `json:"video_introduction,omitempty" xml:"video_introduction,omitempty"`
    // 视频文件地址
    VideoFileUrl   string `json:"video_file_url,omitempty" xml:"video_file_url,omitempty"`
    // 视频预览图
    PreviewUrl   string `json:"preview_url,omitempty" xml:"preview_url,omitempty"`
    // 视频标题
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 视频ID
    VideoId   int64 `json:"video_id,omitempty" xml:"video_id,omitempty"`
}
