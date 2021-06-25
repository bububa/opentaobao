package simba

// SidVo 
type SidVo struct {

    // 当前用户是否有创意本地上传试用功能
    CreativeImgUpload   bool `json:"creative_img_upload,omitempty"`

    // 当前用户是否新版直通车用户
    IsNewBPUser   bool `json:"is_new_b_p_user,omitempty"`

}
