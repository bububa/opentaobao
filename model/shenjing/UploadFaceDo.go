package shenjing

// UploadFaceDo 
type UploadFaceDo struct {
    // 中文消息
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    // 访客总数
    VisitorFacePhotoTotalNum   int64 `json:"visitor_face_photo_total_num,omitempty" xml:"visitor_face_photo_total_num,omitempty"`
    // 出参
    Map   *ResultMap `json:"map,omitempty" xml:"map,omitempty"`
    // extDesc
    ExtDesc   string `json:"ext_desc,omitempty" xml:"ext_desc,omitempty"`
    // 状态码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 当前上传人数
    VisitorFacePhotoCurrentNum   int64 `json:"visitor_face_photo_current_num,omitempty" xml:"visitor_face_photo_current_num,omitempty"`
}
