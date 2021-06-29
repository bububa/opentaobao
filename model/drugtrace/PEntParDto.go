package drugtrace

// PEntParDto 
type PEntParDto struct {
    // 往来单位名称
    PartnerName   string `json:"partner_name,omitempty" xml:"partner_name,omitempty"`
    // 往来单位ID
    PartnerId   string `json:"partner_id,omitempty" xml:"partner_id,omitempty"`
    // 企业id
    EntId   string `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
    // 往来单位企业唯一标识
    RefEntId   string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
    // 往来单位企业所在省编码
    EntProvCode   string `json:"ent_prov_code,omitempty" xml:"ent_prov_code,omitempty"`
    // 所在省
    ProvName   string `json:"prov_name,omitempty" xml:"prov_name,omitempty"`
    // 所在市
    AreaName   string `json:"area_name,omitempty" xml:"area_name,omitempty"`
    // 所在县
    CityName   string `json:"city_name,omitempty" xml:"city_name,omitempty"`
    // 是不是入网企业
    IsNetwork   string `json:"is_network,omitempty" xml:"is_network,omitempty"`
    // 拼音缩写
    PartnerCapitalName   string `json:"partner_capital_name,omitempty" xml:"partner_capital_name,omitempty"`
    // 类型
    PartnerType   string `json:"partner_type,omitempty" xml:"partner_type,omitempty"`
    // 往来单位企业id
    PartnerEntId   string `json:"partner_ent_id,omitempty" xml:"partner_ent_id,omitempty"`
    // 最近修改日期
    LastModDate   string `json:"last_mod_date,omitempty" xml:"last_mod_date,omitempty"`
    // 创建日期
    CrtDate   string `json:"crt_date,omitempty" xml:"crt_date,omitempty"`
    // 创建IC名称
    CrtIcName   string `json:"crt_ic_name,omitempty" xml:"crt_ic_name,omitempty"`
    // 状态
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 修改IC名称
    ModIcName   string `json:"mod_ic_name,omitempty" xml:"mod_ic_name,omitempty"`
    // 级别
    PartnerLevel   string `json:"partner_level,omitempty" xml:"partner_level,omitempty"`
    // 修改IC码
    ModIcCode   string `json:"mod_ic_code,omitempty" xml:"mod_ic_code,omitempty"`
    // 合作ID
    PEntParId   string `json:"p_ent_par_id,omitempty" xml:"p_ent_par_id,omitempty"`
    // 创建IC码
    CrtIcCode   string `json:"crt_ic_code,omitempty" xml:"crt_ic_code,omitempty"`
}
