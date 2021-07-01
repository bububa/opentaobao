package shenjing

// Content 结构体
type Content struct {
	// 访客名称
	VisitorLinkman string `json:"visitor_linkman,omitempty" xml:"visitor_linkman,omitempty"`
	// 访客ID
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 总计来访人数
	VisitorFacePhotoTotalNum int64 `json:"visitor_face_photo_total_num,omitempty" xml:"visitor_face_photo_total_num,omitempty"`
	// 邀请人
	ApplyUserName string `json:"apply_user_name,omitempty" xml:"apply_user_name,omitempty"`
	// 公司名称
	EmpOrgName string `json:"emp_org_name,omitempty" xml:"emp_org_name,omitempty"`
	// 当前上传人脸数
	VisitorFacePhotoCurrentNum int64 `json:"visitor_face_photo_current_num,omitempty" xml:"visitor_face_photo_current_num,omitempty"`
}
