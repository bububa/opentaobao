package vaccin

// VaccineInfo 
type VaccineInfo struct {
    // 疫苗名称
    VaccineName   string `json:"vaccine_name,omitempty" xml:"vaccine_name,omitempty"`
    // 疫苗编码
    VaccineCode   string `json:"vaccine_code,omitempty" xml:"vaccine_code,omitempty"`
    // 疫苗针次
    TheTimes   int64 `json:"the_times,omitempty" xml:"the_times,omitempty"`
    // 阿里健康疫苗编码
    VaccineGbCode   string `json:"vaccine_gb_code,omitempty" xml:"vaccine_gb_code,omitempty"`
}
