package drugtrace

// SeedlingsBaseInfoDTO 
type SeedlingsBaseInfoDTO struct {
    // 基地图片（上传图片）图片建议尺寸：height: 310px;width: 670px;
    BasePictures   []string `json:"base_pictures,omitempty" xml:"base_pictures>string,omitempty"`
    // 品种
    Variety   string `json:"variety,omitempty" xml:"variety,omitempty"`
    // 繁殖部位
    BreedingSite   string `json:"breeding_site,omitempty" xml:"breeding_site,omitempty"`
    // 种苗来源
    SeedlingSource   string `json:"seedling_source,omitempty" xml:"seedling_source,omitempty"`
    // 地理区位
    GeographicLocation   string `json:"geographic_location,omitempty" xml:"geographic_location,omitempty"`
    // 基地面积
    BaseArea   string `json:"base_area,omitempty" xml:"base_area,omitempty"`
    // 基地位置
    BaseLocation   string `json:"base_location,omitempty" xml:"base_location,omitempty"`
    // 基地名称
    BaseName   string `json:"base_name,omitempty" xml:"base_name,omitempty"`
    // 基地认证（文字+图片）图片建议尺寸：height: 310px;width: 670px;
    BaseCertification   *RichTextDTO `json:"base_certification,omitempty" xml:"base_certification,omitempty"`
}
