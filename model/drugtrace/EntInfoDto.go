package drugtrace

// EntInfoDto 
type EntInfoDto struct {

    // 联系方式
    
    EntContact   string `json:"ent_contact,omitempty" xml:"ent_contact,omitempty"`
    

    // 企业资质（上传图片）图片建议尺寸：height: 310px;width: 670px;
    
    EntQualificationPictures   []string `json:"ent_qualification_pictures,omitempty" xml:"ent_qualification_pictures>string,omitempty"`
    

    // 企业介绍
    
    EntIntroduction   string `json:"ent_introduction,omitempty" xml:"ent_introduction,omitempty"`
    

}
